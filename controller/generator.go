package controller

import (
	service_controller "github.com/devingen/devingen-api/controller/service-controller"
	database_service "github.com/devingen/devingen-api/service/database-service"
	kimlik_service "github.com/devingen/kimlik-api/service"
)

// NewServiceController generates new ServiceController
func NewServiceController(service *database_service.DatabaseService, kimlikService kimlik_service.IKimlikService) *service_controller.ServiceController {
	return &service_controller.ServiceController{
		KimlikService: kimlikService,
		Service:       service,
	}
}
