package main

import (
	"github.com/devingen/api-core/database"
	"github.com/devingen/api-core/server"
	"github.com/devingen/devingen-api/controller"
	"github.com/devingen/devingen-api/service"
	"github.com/devingen/kimlik-api/server/wrappers"
	kimlik_service "github.com/devingen/kimlik-api/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Runs the server that contains all the services
func main() {

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Database connection failed %s", err.Error())
	}

	databaseService := service.NewDatabaseService(db)
	kimlikService := kimlik_service.NewDatabaseService(db)
	serviceController := controller.NewServiceController(databaseService, kimlikService)

	router := mux.NewRouter()
	router.HandleFunc("/{base}/workspaces", wrappers.WithLog(wrappers.WithAuth(serviceController.CreateWorkspace))).Methods(http.MethodPost)
	router.HandleFunc("/{base}/workspaces", wrappers.WithLog(wrappers.WithAuth(serviceController.GetWorkspaces))).Methods(http.MethodGet)

	http.Handle("/", &server.CORSRouterDecorator{R: router})
	err = http.ListenAndServe(":1002", &server.CORSRouterDecorator{R: router})
	if err != nil {
		log.Fatalf("Listen and serve failed %s", err.Error())
	}
}
