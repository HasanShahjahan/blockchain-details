package controllers

import (
	"github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/services"
	u "github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/utils"
	v "github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/validation"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	logTag = "Controller"
)

//GetBlockData - gets the details for a specific block.
func GetBlockData(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	network := vars["network"]
	blockhash := vars["blockhash"]

	if v.IsValidNetwork(w, network) {
		return
	}
	if v.IsValidBlockhash(w, blockhash) {
		return
	}

	data, err := services.GetBlock(network, blockhash)
	if err != nil {
		u.ErrorResponse(w, "Failed to get block data", err)
		return
	}

	res := u.Message(true, "Success")
	res["data"] = data
	u.Respond(w, res)
}


