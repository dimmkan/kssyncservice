package main

import (
	"fmt"
	"kssyncservice_go/config"
	"kssyncservice_go/internal/sync"
	"kssyncservice_go/internal/services"
	"kssyncservice_go/pkg/db"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found!")
	}
}

func runSheduller(database *db.Db) {
		// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Println(time.Now(), " - Don't create a new sheduller!")
	}

	syncServicesRepository := sync.NewSyncRepository(database)
	syncSheduller := sync.NewSyncSheduler(sync.SyncShedullerDeps{
		SyncRepository: syncServicesRepository,
	})

	// add a job to the scheduler
	j, err := s.NewJob(
		gocron.CronJob(
			"*/1 * * * *",
			false,
		),
		gocron.NewTask(syncSheduller.SyncRepository.Synchronize, config.New()),
	)
	if err != nil {
		fmt.Println(time.Now(), " - Don't create a new job!")
	}
	// each job has a unique id
	fmt.Println(j.ID())

	// start the scheduler
	s.Start()

	// block until you are ready to shut down
	select {}
}

func main() {
	router := http.NewServeMux()
	database := db.NewDb(config.New())

	//Repositories
	servicesRepository := services.NewServicesRepository(database)

	// Handlers
	services.NewServicesHandler(router, services.ServicesHandlerDeps{
		ServicesRepository: servicesRepository,
	})
	
	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}
	
	fmt.Println("Start sheduller")
	go runSheduller(database)

	fmt.Println("Server started")
	server.ListenAndServe()
}