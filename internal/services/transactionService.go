package services

import (
	"encoding/json"
	logging "github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/logger"
	"github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/models"
	u "github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/utils"
	"time"
)

func GetTransaction(network string, txid string) (*models.TransactionResponse, error) {
	url := "https://sochain.com/api/v2/tx/" + network + "/" + txid
	logging.Info(logTag, "API Request", url)

	responseData, err := u.ResponseWithBytes(url)
	if err != nil {
		return nil, err
	}

	var transactionDTO models.TransactionDTO
	json.Unmarshal(responseData, &transactionDTO)

	var model = models.TransactionResponse{}
	if transactionDTO.Status == "fail" {
		return &model, nil
	}

	model.TXID = transactionDTO.Data.TXID
	model.Fee = transactionDTO.Data.Fee
	model.Sent_Value = transactionDTO.Data.SentValue
	model.Time = time.Unix(transactionDTO.Data.Time, 0).UTC().Format("02/01/2006 15:04")

	return &model, nil
}
