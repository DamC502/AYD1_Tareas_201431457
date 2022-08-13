package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application)routes () http.Handler{
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost,"/suma",app.CalculatorPost)
	return app.enableCORS(router)
}