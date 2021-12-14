package models

//TransactionData - defines our Transaction Data struct
type TransactionData struct {
	TXID      string `json:"txid"`        //Transaction Id
	Fee       string `json:"fee"`         //The transaction fee
	Time      int64  `json:"time"`        //THe transaction Date/Time
	SentValue string `json:"sent_value"`  //The sent value
}

//TransactionDTO - defines our Transaction Data Transfer Object struct
type TransactionDTO struct {
	Data   TransactionData   `json:"data"`
	Status string            `json:"status"`
}

//TransactionResponse - defines our Transaction Response struct
type TransactionResponse struct {
	TXID       string
	Fee        string
	Time       string
	Sent_Value string
}
