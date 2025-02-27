package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"safehouse-main-back/src/internal/controllers"
	mockdb "safehouse-main-back/src/test/mocks/database"
)

func main() {
	log.Println("Starting API in TEST mode with mock database")

	db := mockdb.NewMockDB()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Shutting down gracefully...")
		os.Exit(0)
	}()

	router := controllers.SetupRoutes(db)

	port := ":4000"
	err := router.Run(port)
	if err != nil {
		return
	}
}
