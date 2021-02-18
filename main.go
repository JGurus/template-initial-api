package main

import (
	"log"

	"github.com/JGurus/template-initial-api/api/routes"
	"github.com/JGurus/template-initial-api/auth"
	"github.com/JGurus/template-initial-api/config"
	"github.com/JGurus/template-initial-api/databases"
	"github.com/JGurus/template-initial-api/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := auth.LoadFiles("auth/app.rsa", "auth/app.rsa.pub")
	if err != nil {
		log.Fatalf("certificates could not be loaded %v", err)
	}
	configDB, err := config.GetConfigDB()
	if err != nil {
		log.Fatalf("no se encontróe el archivo de configuración")
	}
	configServer, err := config.GetConfigServer()
	if err != nil {
		log.Fatalf("no se encontróe el archivo de configuración")
	}
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	db := services.NewUser(databases.Factory(configDB.Engine))
	routes.UserRoutes(e, db)
	routes.LoginRoutes(e, db)
	routes.RegisterRoutes(e, db)

	err = e.Start(":" + configServer.Port)
	if err != nil {
		log.Fatal(err)
	}
}
