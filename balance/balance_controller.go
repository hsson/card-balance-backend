// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package balance

import "net/http"
import "fmt"
import "github.com/gorilla/mux"

// GetBalance returns the card balance for the card with the
// specified card number
func GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Get balance for number %v\n", vars["number"])
}
