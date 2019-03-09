package main

import (
	"apparat-api/src/config"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	config.SetupDI()
	router := setupHttpRouter()

	port := config.DI.Configuration.GetPort()
	log.Printf("server listening on %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func setupHttpRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/employees", config.DI.EmployeeController.GetEmployees)
	router.GET("/employees/:name", config.DI.EmployeeController.GetEmployeeByName)
	router.GET("/employeenames", config.DI.EmployeeController.GetEmployeeNames)

	router.ServeFiles("/public/profilepictures/*filepath", http.Dir("public/profilbilder"))

	router.Handle("GET", "/favicon.ico", noopHandler)

	return router
}

func noopHandler(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {}
