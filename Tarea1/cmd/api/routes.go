package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/suma", app.CalculatorPost)
	return app.enableCORS(router)
}

func (app *application) routesInfo() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/info", app.InfoGet)
	return app.enableCORS(router)
}
