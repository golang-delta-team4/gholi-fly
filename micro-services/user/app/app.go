package app

import (
	"context"
	"log"
	"user-service/config"
	"user-service/internal/permission"
	permissionPort "user-service/internal/permission/port"
	"user-service/internal/role"
	rolePort "user-service/internal/role/port"
	"user-service/internal/user"
	userPort "user-service/internal/user/port"
	"user-service/pkg/adapters/clients/grpc"
	"user-service/pkg/adapters/storage"
	"user-service/pkg/adapters/storage/types"
	"user-service/pkg/postgres"

	"gorm.io/gorm"
)

type app struct {
	db                *gorm.DB
	cfg               config.Config
	userService       userPort.Service
	permissionService permissionPort.Service
	roleService       rolePort.Service
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
	a.userService = user.NewService(storage.NewUserRepo(a.db), grpc.NewGRPCBankClient(cfg.BankGRPCConfig.Host, int(cfg.BankGRPCConfig.Port)))
	a.permissionService = permission.NewService(storage.NewPermissionRepo(a.db))
	a.roleService = role.NewService(storage.NewRoleRepo(a.db), a.permissionService, a.userService)
	err := a.roleService.CreateSuperAdminRole(context.Background())
	if err != nil {
		log.Println(err)
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

func (a *app) UserService() userPort.Service {
	return a.userService
}

func (a *app) PermissionService() permissionPort.Service {
	return a.permissionService
}

func (a *app) RoleService() rolePort.Service {
	return a.roleService
}
