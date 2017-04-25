// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package menu

import (
	"fmt"
	"net/http"
)

// Index gets the entire food menu
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get the food menu")
}
