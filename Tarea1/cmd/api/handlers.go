package main

import (
	"encoding/json"
	"net/http"
)

type CalculatorReq struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

func (app *application) CalculatorPost(w http.ResponseWriter, r *http.Request) {
	var req CalculatorReq
	errReq := json.NewDecoder(r.Body).Decode(&req) //json -> parser -> struct
	if errReq != nil {
		app.logger.Println(errReq)
		errorJSON(w, errReq)
		return
	}
	type CalculatorRes struct {
		Result int `json:"result"`
	}

	result := CalculatorRes{Result: req.Number1 + req.Number2}

	err := writeJSON(w, http.StatusOK, result, "")

	if err != nil {
		errorJSON(w, err)
	}
	return
}

func (app *application) SubPost(w http.ResponseWriter, r *http.Request) {
	var req CalculatorReq
	errReq := json.NewDecoder(r.Body).Decode(&req) //json -> parser -> struct
	if errReq != nil {
		app.logger.Println(errReq)
		errorJSON(w, errReq)
		return
	}
	type CalculatorRes struct {
		Result int `json:"result"`
	}

	result := CalculatorRes{Result: req.Number1 - req.Number2}

	err := writeJSON(w, http.StatusOK, result, "")

	if err != nil {
		errorJSON(w, err)
	}
	return
}

func (app *application) InfoGet(w http.ResponseWriter, r *http.Request) {
	type InfoRes struct {
		Name   string `json:"name"`
		Carnet int    `json:"carnet"`
	}

	res := InfoRes{Name: "Damihan Morales", Carnet: 201431457}

	err := writeJSON(w, http.StatusOK, res, "datos")
	if err != nil {
		errorJSON(w, err)
	}
	return

}
