package app

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/config"
	companyPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/port"
	technicalTeam "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/technicalTeam/port"
	ticketPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket/port"
	tripPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/port"

	"gorm.io/gorm"
)

type App interface {
	CompanyService(ctx context.Context) companyPort.Service
	TripService(ctx context.Context) tripPort.Service
	TicketService(ctx context.Context) ticketPort.Service
	TechnicalTeamService(ctx context.Context) technicalTeam.Service
	DB() *gorm.DB
	Config() config.Config
}
