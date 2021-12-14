package utils

import (
	"encoding/json"
	logging "github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/logger"
	"io/ioutil"
	"net/http"
)

const (
	logTag = "Utils"
)
//Response - stores response from API
type Response struct {
	Message string
	Error   string
}

//Message ...
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Respond ...
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

//ErrorResponse ...
func ErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}

//RespondWithJSON ...
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//RespondWithError ...
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

//ResponseWithBytes ...
func ResponseWithBytes(url string)  ([]byte, error){
	response, err := http.Get(url)
	if err != nil {
		logging.Error(logTag, "error getting response, error=%+v", err)
		return nil, err
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logging.Error(logTag, "error reading response, error=%+v", err)
		return nil, err
	}

	logging.Info(logTag, "API Response: \n", string(responseData))
	return responseData, nil
}