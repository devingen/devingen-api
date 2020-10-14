module github.com/devingen/devingen-api

go 1.12

//replace github.com/devingen/api-core => ../api-core

//replace github.com/devingen/kimlik-api => ../kimlik-api

require (
	github.com/devingen/api-core v0.0.13
	github.com/devingen/kimlik-api v0.0.6
	github.com/gorilla/mux v1.7.4
	go.mongodb.org/mongo-driver v1.3.2
)
