// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package menu

import "net/http"

// Index gets the entire food menu
func Index(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "Get the food menu", nil
}
