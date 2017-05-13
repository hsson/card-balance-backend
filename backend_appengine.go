// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build appengine

package backend

import "net/http"

// Run starts the API for appengine
func Run() {
	r := newRouter()
	http.Handle("/", r)
}
