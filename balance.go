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
	"io/ioutil"
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

func fetchTransactions(address string) ([]TransactionBlock, int, error) {

	// Construire l'URL de l'API
	url := fmt.Sprintf("http://192.168.1.8:3000/account/%s/history/10000000", address)

	// Appel HTTP à l'API
	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, err
	}

	// Fermer le corps de réponse après la fonction
	defer resp.Body.Close()

	// Lire le corps de réponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	// Décoder la réponse JSON
	var history HistoryResponse
	if err := json.Unmarshal(body, &history); err != nil {
		return nil, 0, err
	}

	// Retourner les données
	return history.Data.PerBlock, history.Code, nil
}
