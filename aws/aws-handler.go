package aws

import (
	"github.com/devingen/devingen-api/controller"
	service_controller "github.com/devingen/devingen-api/controller/service-controller"
	"github.com/devingen/devingen-api/service"
	kimlik_service "github.com/devingen/kimlik-api/service"
)

func GenerateController() *service_controller.ServiceController {
	db := GetDatabase()
	databaseService := service.NewDatabaseService(db)
	kimlikService := kimlik_service.NewDatabaseService(db)
	return controller.NewServiceController(databaseService, kimlikService)
}
