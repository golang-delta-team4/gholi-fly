package http

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/api/service"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/app"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/config"
)

// company service transient instance handler
func companyServiceGetter(appContainer app.App, cfg config.ServerConfig) ServiceGetter[*service.CompanyService] {
	return func(ctx context.Context) *service.CompanyService {
		return service.NewCompanyService(appContainer.CompanyService(ctx))
	}
}
