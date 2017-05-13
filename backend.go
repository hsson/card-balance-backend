// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package backend

import "net/http"

func Run() {
	r := newRouter()
	http.Handle("/", r)
}
