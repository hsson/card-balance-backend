// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package balance

import (
	"net/http"

	"errors"

	"github.com/gorilla/mux"
	"github.com/hsson/card-balance-backend/backend/modules"
	backendConfig "github.com/hsson/card-balance-backend/config"
)

const allowedCardNumberLength = 16

// ErrorInvalidCardNumber is sent when the card number is not 16 digits long
var ErrorInvalidCardNumber = errors.New("Invalid card number")

var config backendConfig.Config

// Data represents the balance data available from the card website
type Data struct {
	CardNumber string  `json:"card_number"`
	FullName   string  `json:"full_name"`
	Email      string  `json:"email"`
	Balance    float64 `json:"balance"`
}

// Init initializes the module with given config
func Init(newConfig backendConfig.Config) {
	config = newConfig
}

// GetBalance returns the card balance for the card with the
// specified card number
func GetBalance(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	cardNumber := vars["number"]
	if len(cardNumber) != allowedCardNumberLength {
		return nil, ErrorInvalidCardNumber
	}
	scraper := new(scraper)
	// Initialize scraper for new scrape request
	scraper.init()
	scraper.client = modules.GetHTTPClient(r)
	data, err := scraper.Scrape(cardNumber)
	if err != nil {
		return nil, err
	}
	return data, nil
}
