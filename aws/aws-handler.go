package aws

import (
	"github.com/devingen/devingen-api/controller"
	service_controller "github.com/devingen/devingen-api/controller/service-controller"
	"github.com/devingen/devingen-api/service"
)

func GenerateController() *service_controller.ServiceController {
	db := GetDatabase()
	databaseService := service.NewDatabaseService(db)
	return controller.NewServiceController(databaseService)
}
