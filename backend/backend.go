// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build !appengine

package backend

import (
	"fmt"
	"net/http"

	"github.com/hsson/card-balance-backend/backend/modules/balance"
	"github.com/hsson/card-balance-backend/backend/modules/charge"
	"github.com/hsson/card-balance-backend/backend/modules/menu"
	backendConfig "github.com/hsson/card-balance-backend/config"
)

const (
	configFile = "./config/config.yaml"
)

// Run starts the API for appengine
func Run() {
	conf := backendConfig.LoadConfig(configFile)
	fmt.Printf("Config loaded, URL is %v\n", conf.CardSiteURL)
	menu.Init(conf)
	charge.Init(conf)
	balance.Init(conf)

	r := newRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
