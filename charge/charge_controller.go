// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package charge

import "net/http"

// TODO: Move to config file
const chargingWebsiteURL = "https://kortladdning3.chalmerskonferens.se/"

// Redirect redirects to the proper website for charging a card
func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, chargingWebsiteURL, http.StatusTemporaryRedirect)
}
