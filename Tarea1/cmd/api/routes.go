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
	router.HandlerFunc(http.MethodPost, "/resta", app.SubPost)
	return app.enableCORS(router)
}
