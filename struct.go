/*
*****************************************************************************
*
*                                Warthog Explorer
*
*****************************************************************************
* Web : https://Makertronic-yt.com
* copyright 2023
 */

package main

type BalanceData struct {
	AccountID int    `json:"accountId"`
	Balance   string `json:"balance"`
}

type HashRateData struct {
	Last100BlocksEstimate int `json:"last100BlocksEstimate"`
}

type ChainHeadData struct {
	Height     int     `json:"height"`
	Difficulty float64 `json:"difficulty"`
}

// Define structs matching the JSON structure
type TransactionBlock struct {
	Confirmations int `json:"confirmations"`
	Height        int `json:"height"`
	Transactions  struct {
		Rewards []struct {
			Amount    string `json:"amount"`
			ToAddress string `json:"toAddress"`
			TxHash    string `json:"txHash"`
		} `json:"rewards"`
	} `json:"transactions"`
}

type HistoryResponse struct {
	Code int `json:"code"`
	Data struct {
		PerBlock []TransactionBlock `json:"perBlock"`
	} `json:"data"`
}

type PageData struct {
	Version         string
	Balance         string
	HashRate        int
	ChainHeight     int
	ChainDifficulty float64
	Transactions    []TransactionBlock
}
