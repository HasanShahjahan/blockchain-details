package models

//BlockData - defines our Block Data struct
type BlockData struct {
	Network           string   `json:"network"`              //The acronym of the network.
	BlockNo           int64    `json:"block_no"`             //The height of the block in the blockchain, or its number.
	Time              int64    `json:"time"`                 //The time at which this block was mined by the miner.
	PreviousBlockHash string   `json:"previous_blockhash"`   //The block hash of the previous block in the blockchain.
	NextBlockHash     string   `json:"next_blockhash"`       //The block hash of the next block in the blockchain. next_blockhash=null if this is the last block in the blockchain.
	Size              int      `json:"size"`                 //The size of the block in bytes.
	Transactions      []string `json:"txs"`
}

//BlockDTO - defines our Block Data Transfer Object struct
type BlockDTO struct {
	Data   BlockData   `json:"data"`
	Status string      `json:"status"`
}

//BlockResponse - defines our Block Response struct
type BlockResponse struct {
	Network             string
	BlockNo             int64
	Time                string
	Previous_Block_Hash string
	Next_Block_Hash     string
	Size                int
	Transactions        [10]TransactionResponse
}
