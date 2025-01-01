package app

import (
	"context"
	"fmt"
	"log"

	// "log"
	clientPort "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/port"
	clientHttp "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/http"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/types"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/config"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company"
	companyPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/invoice"
	invoicePort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/invoice/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/technicalTeam"
	technicalTeamPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/technicalTeam/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket"
	ticketPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip"
	tripPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage"

	// "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/types"

	// "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/types"
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
	tripService    tripPort.Service
	ticketService  ticketPort.Service
	technicalTeam  technicalTeamPort.Service
	redisProvider  cache.Provider
	userGRPCClient clientPort.GRPCUserClient
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
	return company.NewService(storage.NewCompanyRepo(db, false, a.redisProvider), grpc.NewGRPCRoleClient(a.cfg.Role.Host, int(a.cfg.Role.Port)))
}

func (a *app) TripService(ctx context.Context) tripPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.tripService == nil {
			a.tripService = a.tripServiceWithDB(a.db)
		}
		return a.tripService
	}

	return a.tripServiceWithDB(db)
}

func (a *app) tripServiceWithDB(db *gorm.DB) tripPort.Service {
	return trip.NewService(
		storage.NewTripRepo(db, false, a.redisProvider),
		storage.NewTechnicalTeamRepo(db, false, a.redisProvider),
		storage.NewTripRepo(db, false, a.redisProvider),
		clientHttp.NewHttpPathClient(int(a.cfg.Map.Port)),
		clientHttp.NewHttpVehicleClient(int(a.cfg.Vehicle.Port)),
		a.CompanyService(context.Background()))
}

func (a *app) TicketService(ctx context.Context) ticketPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.ticketService == nil {
			a.ticketService = a.ticketServiceWithDB(a.db)
		}
		return a.ticketService
	}

	return a.ticketServiceWithDB(db)
}

func (a *app) ticketServiceWithDB(db *gorm.DB) ticketPort.Service {
	return ticket.NewService(storage.NewTicketRepo(db, false, a.redisProvider),
		a.tripServiceWithDB(db), a.invoiceServiceWithDB(db), grpc.NewGRPCBankClient(a.cfg.Bank.Host, int(a.cfg.Bank.Port)))
}

func (a *app) TechnicalTeamService(ctx context.Context) technicalTeamPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.technicalTeam == nil {
			a.technicalTeam = a.technicalTeamServiceWithDB(a.db)
		}
		return a.technicalTeam
	}

	return a.technicalTeamServiceWithDB(db)
}

func (a *app) technicalTeamServiceWithDB(db *gorm.DB) technicalTeamPort.Service {
	return technicalTeam.NewService(storage.NewTechnicalTeamRepo(db, false, a.redisProvider))
}

func (a *app) invoiceServiceWithDB(db *gorm.DB) invoicePort.Service {
	return invoice.NewService(storage.NewInvoiceRepo(db, false, a.redisProvider))
}

func (a *app) UserGRPCService() clientPort.GRPCUserClient {
	return a.userGRPCClient
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

	migrateErr := db.AutoMigrate(&types.Company{}, &types.Ticket{}, &types.Invoice{}, &types.TechnicalTeam{}, &types.TechnicalTeamMember{}, &types.Trip{}, &types.VehicleRequest{})
	if migrateErr != nil {
		log.Fatalf("Failed to migrate : %v", migrateErr)
	}

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
	a.userGRPCClient = grpc.NewGRPCUserClient(a.cfg.User.Host, int(a.cfg.User.Port))

	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
