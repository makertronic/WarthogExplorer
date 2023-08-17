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

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BalanceResponse struct {
	Code int         `json:"code"`
	Data BalanceData `json:"data"`
}

func fetchBalance(address string) (BalanceData, error) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	fmt.Print("fetchBalance : ", formattedTime, " - ", address)
	apiURL := fmt.Sprintf("http://%s:%s/account/%s/balance", NODE, NODE_PORT, address)
	resp, err := http.Get(apiURL)
	if err != nil {
		return BalanceData{}, err
	}
	defer resp.Body.Close()

	var balanceResponse BalanceResponse

	err = json.NewDecoder(resp.Body).Decode(&balanceResponse)
	if err != nil {
		fmt.Println("Error : ", err)
		return BalanceData{}, err
	}

	fmt.Println(" - ", balanceResponse.Data)
	return balanceResponse.Data, nil
}
