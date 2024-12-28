package app

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/config"
	companyPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/port"
	ticketPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket/port"
	tripPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/port"
	clientPort "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/port"

	"gorm.io/gorm"
)

type App interface {
	CompanyService(ctx context.Context) companyPort.Service
	TripService(ctx context.Context) tripPort.Service
	TicketService(ctx context.Context) ticketPort.Service
	UserGRPCService() clientPort.GRPCUserClient
	DB() *gorm.DB
	Config() config.Config
}
