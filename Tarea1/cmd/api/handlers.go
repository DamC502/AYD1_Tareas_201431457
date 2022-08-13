package main

import (
	"encoding/json"
	"net/http"
)

type CalculatorReq struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

func (app *application) CalculatorPost( w http.ResponseWriter, r *http.Request ) {
	var req CalculatorReq
	errReq := json.NewDecoder(r.Body).Decode(&req)  //json -> parser -> struct
	if errReq != nil {
		app.logger.Println(errReq)
		errorJSON(w, errReq)
		return
	}
	type CalculatorRes struct {
		Result int `json:"result"`
	}

	result := CalculatorRes{ Result: req.Number1 + req.Number2}

	err := writeJSON(w, http.StatusOK, result, "")

	if err != nil {
		errorJSON(w, err)
	}
	return
}

