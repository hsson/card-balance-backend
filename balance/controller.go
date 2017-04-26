// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package balance

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const allowedCardNumberLength = 16

// Data represents the balance data available from the card website
type Data struct {
	CardNumber string  `json:"card_number"`
	FullName   string  `json:"full_name"`
	Email      string  `json:"email"`
	Balance    float64 `json:"balance"`
}

// GetBalance returns the card balance for the card with the
// specified card number
func GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardNumber := vars["number"]
	scraper := new(scraper)
	data, err := scraper.Scrape(cardNumber)
	if err != nil {
		// TODO: Proper error handling
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
