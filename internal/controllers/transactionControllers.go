package controllers

import (
	"github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/services"
	u "github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/utils"
	v "github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/validation"
	"net/http"

	"github.com/gorilla/mux"
)

//GetTransactionData - gets the details for a specific transaction.
func GetTransactionData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	network := vars["network"]
	txid := vars["txid"]

	if v.IsValidNetwork(w, network) {
		return
	}
	if v.IsValidTransactionId(w, txid) {
		return
	}

	data, err := services.GetTransaction(network, txid)
	if err != nil {
		u.ErrorResponse(w, "Failed to get transaction data", err)
		return
	}

	res := u.Message(true, "SUCCESS!")
	res["data"] = data
	u.Respond(w, res)
}
