// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package charge

import (
	"net/http"

	backendConfig "github.com/hsson/card-balance-backend/config"
)

var config backendConfig.Config

// Init intializes the module with given config
func Init(newConfig backendConfig.Config) {
	config = newConfig
}

// Redirect redirects to the proper website for charging a card
func Redirect(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	http.Redirect(w, r, config.CardSiteURL, http.StatusTemporaryRedirect)
	return nil, nil
}
