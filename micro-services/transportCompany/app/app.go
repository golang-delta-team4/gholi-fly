package app

import (
	"context"
	"fmt"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/config"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company"
	companyPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/types"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/cache"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/postgres"

	redisAdapter "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/cache"

	"gorm.io/gorm"

	appCtx "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/context"
)

type app struct {
	db             *gorm.DB
	cfg            config.Config
	companyService companyPort.Service
	redisProvider  cache.Provider
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) CompanyService(ctx context.Context) companyPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.companyService == nil {
			a.companyService = a.companyServiceWithDB(a.db)
		}
		return a.companyService
	}

	return a.companyServiceWithDB(db)
}

func (a *app) companyServiceWithDB(db *gorm.DB) companyPort.Service {
	return company.NewService(storage.NewCompanyRepo(db, false, a.redisProvider))
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

	db.AutoMigrate(&types.Company{}, &types.Invoice{}, &types.TechnicalTeam{}, &types.TechnicalTeamMemeber{}, &types.Ticket{}, &types.Ticket{}, &types.Trip{}, &types.VehicleRequest{})

	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *app) setRedis() {
	a.redisProvider = redisAdapter.NewRedisProvider(fmt.Sprintf("%s:%d", a.cfg.Redis.Host, a.cfg.Redis.Port))
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	a.setRedis()

	//return a, a.registerOutboxHandlers()

	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}

// func (a *app) registerOutboxHandlers() error {
// 	scheduler, err := gocron.NewScheduler()
// 	if err != nil {
// 		return err
// 	}

// 	common.RegisterOutboxRunner(a.notifServiceWithDB(a.db), scheduler)

// 	scheduler.Start()

// 	return nil
// }
