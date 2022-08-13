package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const version = "1"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status     string `json:"status"`
	Enviroment string `json:"enviroment"`
	Version    string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	var cfg2 config
	cfg = config{
		port: 3000,
		env:  "development",
	}

	cfg2 = config{
		port: 4000,
		env:  "development",
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server2 := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg2.port),
		Handler:      app.routesInfo(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		logger.Println("Starting server on port", cfg.port)
		err := server.ListenAndServe()

		if err != nil {
			log.Println(err)
		}
		wg.Done()
	}()

	go func() {
		logger.Println("Starting server on port", cfg2.port)
		err := server2.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
		wg.Done()
	}()

	wg.Wait()

}
