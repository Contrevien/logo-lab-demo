package main

import (
	"log"
	"time"

	"logo-lab-demo/bootstrap"
	"logo-lab-demo/middleware"
	route "logo-lab-demo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	port := env.ServerAddress
	if port == "" {
		port = "8080"
	}

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	route.Setup(env, timeout, db, server)

	if err := server.Run(":" + port); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
