// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import "net/http"

func main() {
	r := NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
