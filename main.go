/*
*****************************************************************************
*
*                              Warthog Explorer
*
*****************************************************************************
* Web : https://Makertronic-yt.com
* copyright 2023
 */

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/klauspost/lctime"
	"github.com/logrusorgru/aurora/v4"
)

const (
	VERSION   = "1.0.1"
	NODE      = "192.168.1.8"
	NODE_PORT = "3000"
	HTTP_PORT = "8080"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
}

func handlePOST(w http.ResponseWriter, r *http.Request) {

	// Récupération des données
	address := r.FormValue("address")

	balance, err := fetchBalance(address)
	if err != nil {
		http.Error(w, "Failed to fetch balance", http.StatusInternalServerError)
		return
	}

	// Idem pour hashrate, chainHead, transactions

	transactions, code, err := fetchTransactions(address)
	if err != nil {
		http.Error(w, "Failed to fetch transactions", http.StatusInternalServerError)
		return
	}

	if code != 0 {
		log.Printf("Transactions API error: %d", code)
	}

	// Construction de la réponse
	data := PageData{
		Balance:      balance.Balance,
		Transactions: transactions,
		Version:      VERSION,
	}

	data.Balance = fmt.Sprintf("Account ID: %d, Balance: %s", balance.AccountID, balance.Balance)

	// Affichage de la réponse
	tmpl.Execute(w, data)

}

func handleGET(w http.ResponseWriter, r *http.Request) {

	var data PageData

	// Chargement initial de la page
	chainHead, err := fetchChainHead()
	if err != nil {
		http.Error(w, "Failed to fetch chain head data", http.StatusInternalServerError)
		return
	}

	data.ChainHeight = chainHead.Height
	data.ChainDifficulty = chainHead.Difficulty

	hashRate, err := fetchHashRate()
	if err != nil {
		http.Error(w, "Failed to fetch hash rate", http.StatusInternalServerError)
		return
	}
	data.HashRate = hashRate.Last100BlocksEstimate
	// Rendu du template
	tmpl.Execute(w, data)
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		handlePOST(w, r)
	} else {
		handleGET(w, r)
	}

}

func main() {
	// show something at startup  ------------------------------------------------------------------------------------------

	fmt.Print("\033[H\033[2J") // clear screen
	fmt.Println(aurora.Bold(aurora.Blue("░▒█░░▒█░█▀▀▄░█▀▀▄░▀█▀░█░░░░▄▀▀▄░█▀▀▀░░░▒█▀▀▀░█░█░▄▀▀▄░█░░▄▀▀▄░█▀▀▄░█▀▀░█▀▀▄")))
	fmt.Println(aurora.Bold(aurora.White("░▒█▒█▒█░█▄▄█░█▄▄▀░░█░░█▀▀█░█░░█░█░▀▄░░░▒█▀▀▀░▄▀▄░█▄▄█░█░░█░░█░█▄▄▀░█▀▀░█▄▄▀")))
	fmt.Println(aurora.Bold(aurora.Cyan("░▒▀▄▀▄▀░▀░░▀░▀░▀▀░░▀░░▀░░▀░░▀▀░░▀▀▀▀░░░▒█▄▄▄░▀░▀░█░░░░▀▀░░▀▀░░▀░▀▀░▀▀▀░▀░▀▀")))
	fmt.Printf("░> A Warthog Explorer v%s - Makertronic 2023\n", aurora.Green(VERSION))
	fmt.Println(lctime.Strftime("░> Use help flag for help - Started at %c", time.Now()))
	fmt.Println("░> Warthog Explorer is running on port", aurora.Green(HTTP_PORT))
	fmt.Println(" ")

	fmt.Println(aurora.Cyan("░> Logs:"))

	http.HandleFunc("/", handleHomePage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	port := fmt.Sprintf(":%s", HTTP_PORT)
	fmt.Printf("Server listening on port %s...\n", HTTP_PORT)
	http.ListenAndServe(port, nil)

}
