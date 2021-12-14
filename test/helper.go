package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

type ResponseModel struct {
	Status bool `json:"status"`
}

func ApiHelper(t *testing.T, url string) bool {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fail()
	}

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fail()
	}

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fail()
	}

	blockStruct := ResponseModel{}
	json.Unmarshal(response, &blockStruct)
	return blockStruct.Status
}
