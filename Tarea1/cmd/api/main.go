package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	cfg = config{
		port: 3000,
		env:  "development",
	}
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout:  10* time.Second,
		WriteTimeout: 10* time.Second,
	}


	logger.Println("Starting server on port", cfg.port)
	err := server.ListenAndServe()

	if err != nil {
		log.Println(err)
	}
	return
}
