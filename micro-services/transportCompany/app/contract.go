package app

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/config"
	companyPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/company/port"

	"gorm.io/gorm"
)

type App interface {
	CompanyService(ctx context.Context) companyPort.Service
	DB() *gorm.DB
	Config() config.Config
}
