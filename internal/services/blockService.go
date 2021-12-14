package services

import (
	"encoding/json"
	"github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/logger"
	"github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/models"
	u "github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/utils"
	"time"
)

const (
	logTag = "Service"
)

func GetBlock(network string, blockhash string) (*models.BlockResponse, error) {
	url := "https://sochain.com/api/v2/get_block/" + network + "/" + blockhash
	logging.Info(logTag, "API Request", url)

	responseData, err := u.ResponseWithBytes(url)
	if err != nil {
		return nil, err
	}

	//Convert response body to DTO
	blockDTO := models.BlockDTO{}
	json.Unmarshal(responseData, &blockDTO)

	var model = models.BlockResponse{}
	if blockDTO.Status == "fail" {
		return &model, nil
	}

	model.Network = blockDTO.Data.Network
	model.BlockNo = blockDTO.Data.BlockNo
	model.Next_Block_Hash = blockDTO.Data.NextBlockHash
	model.Previous_Block_Hash = blockDTO.Data.PreviousBlockHash
	model.Size = blockDTO.Data.Size
	model.Time = time.Unix(blockDTO.Data.Time, 0).UTC().Format("02/01/2006 15:04")

	for i := 0; i < 10; i++ {
		model.Transactions[i].TXID = blockDTO.Data.Transactions[i]
		data, err := GetTransaction(network, blockDTO.Data.Transactions[i])
		if err == nil {
			model.Transactions[i].Fee = data.Fee
			model.Transactions[i].Sent_Value = data.Sent_Value
			model.Transactions[i].Time = data.Time
		}
	}
	return &model, nil
}
