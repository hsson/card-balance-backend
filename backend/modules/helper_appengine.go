// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build appengine

package modules

import (
	"errors"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

var errorNoRedirect = errors.New("Don't redirect")

func RedirectFunc(req *http.Request, via []*http.Request) error {
	return errorNoRedirect
}

func GetHTTPClient(r *http.Request) *http.Client {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	client.Timeout = time.Second * 10
	return client
}
