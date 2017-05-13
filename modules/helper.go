// +build !appengine

package modules

import (
	"net/http"
	"time"
)

func RedirectFunc(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func GetHTTPClient(r *http.Request) *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}
