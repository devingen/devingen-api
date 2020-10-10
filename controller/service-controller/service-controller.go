package service_controller

import (
	"github.com/devingen/devingen-api/service"
	kimlik_service "github.com/devingen/kimlik-api/service"
)

// ServiceController implements IServiceController interface by using IDevingenService
type ServiceController struct {
	Service       service.IDevingenService
	KimlikService kimlik_service.IKimlikService
}
