package app

import (
	"context"
	"log"
	"os"
	"user-service/config"
	"user-service/internal/permission"
	"user-service/internal/permission/domain"
	permissionPort "user-service/internal/permission/port"
	"user-service/internal/role"
	roleDomain "user-service/internal/role/domain"
	rolePort "user-service/internal/role/port"
	"user-service/internal/user"
	userDomain "user-service/internal/user/domain"
	userPort "user-service/internal/user/port"
	"user-service/pkg/adapters/clients/grpc"
	grpcPort "user-service/pkg/adapters/clients/grpc/port"
	"user-service/pkg/adapters/storage"
	"user-service/pkg/adapters/storage/types"
	appCtx "user-service/pkg/context"
	"user-service/pkg/postgres"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type app struct {
	db                *gorm.DB
	cfg               config.Config
	userService       userPort.Service
	permissionService permissionPort.Service
	roleService       rolePort.Service
	bankClient        grpcPort.GRPCBankClient
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Password,
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		DBName: a.cfg.DB.Database,
		Schema: a.cfg.DB.Schema,
	})

	if err != nil {
		return err
	}

	a.db = db

	err = autoMigrate(db)

	if err != nil {
		return err
	}

	return nil
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&types.User{},
		&types.RefreshToken{},
		types.Permission{},
		&types.Role{},
		&types.UserRole{})
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}
	a.bankClient = grpc.NewGRPCBankClient(cfg.BankGRPCConfig.Host, int(cfg.BankGRPCConfig.Port))
	a.userService = user.NewService(storage.NewUserRepo(a.db), grpc.NewGRPCBankClient(cfg.BankGRPCConfig.Host, int(cfg.BankGRPCConfig.Port)))
	a.permissionService = permission.NewService(storage.NewPermissionRepo(a.db))
	a.roleService = role.NewService(storage.NewRoleRepo(a.db), a.permissionService, a.userService)
	roleUUID, err := a.roleService.CreateSuperAdminRole(context.Background())
	if err != nil {
		log.Printf("error creating super admin: %v\n", err)
	}
	_, err = a.permissionService.CreatePermissions(context.Background(), seedPermissions())
	if err != nil {
		log.Printf("error seeding permissions: %v\n", err)
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	userID, err := a.userService.SignUp(context.Background(), &userDomain.User{
		FirstName: os.Getenv("DEFAULT_USER_FIRST_NAME"),
		LastName:  os.Getenv("DEFAULT_USER_LAST_NAME"),
		Email:     os.Getenv("DEFAULT_USER_EMAIL"),
		Password:  os.Getenv("DEFAULT_USER_PASSWORD")})

	if err != nil {
		log.Printf("error creating user: %v\n", err)
	}
	err = a.roleService.AssignRole(context.Background(), userID, []roleDomain.Role{{UUID: roleUUID}})
	if err != nil {
		log.Printf("error assigning roles: %v\n", err)
	}
	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}

func (a *app) roleServiceWithDB(db *gorm.DB) rolePort.Service {
	return role.NewService(storage.NewRoleRepo(db), a.permissionService, a.userService)
}

func (a *app) RoleService(ctx context.Context) rolePort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.roleService == nil {
			a.roleService = a.roleServiceWithDB(a.db)
		}
		return a.roleService
	}

	return a.roleServiceWithDB(db)
}

func (a *app) userServiceWithDB(db *gorm.DB) userPort.Service {
	return user.NewService(storage.NewUserRepo(db), a.bankClient)
}

func (a *app) UserService(ctx context.Context) userPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.userService == nil {
			a.userService = a.userServiceWithDB(a.db)
		}
		return a.userService
	}

	return a.userServiceWithDB(db)
}

func (a *app) permissionServiceWithDB(db *gorm.DB) permissionPort.Service {
	return permission.NewService(storage.NewPermissionRepo(db))
}

func (a *app) PermissionService(ctx context.Context) permissionPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.permissionService == nil {
			a.permissionService = a.permissionServiceWithDB(a.db)
		}
		return a.permissionService
	}

	return a.permissionServiceWithDB(db)
}

func seedPermissions() []domain.Permission {
	permissions := []domain.Permission{
		{
			Route:  "/users",
			Method: "GET",
		},
		{
			Route:  "/permissions",
			Method: "POST",
		},
		{
			Route:  "/roles",
			Method: "POST",
		},
		{
			Route:  "/roles",
			Method: "GET",
		},
		{
			Route:  "/roles/assign",
			Method: "POST",
		},
		{
			Route:  "/roles/:id",
			Method: "DELETE",
		},
	}

	return permissions
}
