package main

import (
	"github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/controllers"
	logging "github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/logger"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	logTag = "Start"
)

func main() {
	logging.SetLogLevel("INFO")

	router := mux.NewRouter()
	router.HandleFunc("/api/transaction/{network}/{txid}", controllers.GetTransactionData).Methods("GET")
	router.HandleFunc("/api/block/{network}/{blockhash}", controllers.GetBlockData).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	logging.Info(logTag, "Port", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		logging.Fatal(logTag, "error getting start, error=%+v", err)
	}
}