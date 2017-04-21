// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package charge

import "net/http"
import "fmt"

// Redirect redirects to the proper website for
// charging a card
func Redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Should redirect to the proper website")
}
