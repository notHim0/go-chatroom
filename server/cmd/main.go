package main

import (
	"log"
	"os"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
	"server/router"

	"github.com/joho/godotenv"
)

func main(){
	var err error = godotenv.Load()

	//checking for any env variable loading erors
	if err != nil {
		log.Fatalf("Error loading env file: %s",err.Error())
	}

	var connectionString string = os.Getenv("DATABASE_URI")
	if len(connectionString) == 0 {
		log.Fatalf("Database uri is not set")
	}
	dbConn, err := db.NewDatabase(connectionString)

	if err!=nil {
		log.Fatalf("Could not initailise db connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}