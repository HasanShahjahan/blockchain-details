package validation

import (
	u "github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/utils"
	"net/http"
)

//IsValidNetwork - validates only the following cryptocurrency network codes: BTC, LTC and DOGE.
func IsValidNetwork(w http.ResponseWriter, network string) bool {
	if network != "BTC" && network != "LTC" && network != "DOGE" {
		w.WriteHeader(http.StatusBadRequest)
		res := u.Message(false, "NETWORK NOT SUPPORTED")
		u.Respond(w, res)
		return true
	}
	return false
}

//IsValidBlockhash - validates reference number for a block in the blockchain which length is 64 (SHA256 twice).
func IsValidBlockhash(w http.ResponseWriter, blockhash string) bool {
	if len(blockhash) != 64 {
		w.WriteHeader(http.StatusBadRequest)
		res := u.Message(false, "INVALID BLOCKHASH")
		u.Respond(w, res)
		return true
	}
	return false
}

//IsValidTransactionId - validates transaction id which length is 64 (SHA256 twice).
func IsValidTransactionId(w http.ResponseWriter, txid string) bool {
	if len(txid) != 64 {
		w.WriteHeader(http.StatusBadRequest)
		res := u.Message(false, "INVALID TXID")
		u.Respond(w, res)
		return true
	}
	return false
}


