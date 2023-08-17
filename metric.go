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
)

// network hash ----------------------------------------------------------
func fetchHashRate() (HashRateData, error) {
	url := fmt.Sprintf("http://%s:%s/chain/hashrate", NODE, NODE_PORT)

	// Effectuer une requête GET à l'URL
	response, err := http.Get(url)
	if err != nil {
		return HashRateData{}, err
	}
	defer response.Body.Close()

	// Lire et décoder la réponse JSON
	var responseData struct {
		Code int          `json:"code"`
		Data HashRateData `json:"data"`
	}
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		return HashRateData{}, err
	}

	// Vérifier le code de réponse
	if responseData.Code != 0 {
		return HashRateData{}, fmt.Errorf("API returned an error: code %d", responseData.Code)
	}

	return responseData.Data, nil
}

// height& difficulty ----------------------------------------------------------
func fetchChainHead() (ChainHeadData, error) {
	url := fmt.Sprintf("http://%s:%s/chain/head", NODE, NODE_PORT)

	response, err := http.Get(url)
	if err != nil {
		return ChainHeadData{}, err
	}
	defer response.Body.Close()

	var responseData struct {
		Code int           `json:"code"`
		Data ChainHeadData `json:"data"`
	}
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		return ChainHeadData{}, err
	}

	if responseData.Code != 0 {
		return ChainHeadData{}, fmt.Errorf("API returned an error: code %d", responseData.Code)
	}

	return responseData.Data, nil
}
