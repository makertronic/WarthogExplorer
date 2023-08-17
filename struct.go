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

type PageData struct {
	Balance         string
	HashRate        int
	ChainHeight     int
	ChainDifficulty float64
}
