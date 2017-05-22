// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build appengine

package backend

import (
	"net/http"

	"github.com/hsson/card-balance-backend/backend/modules/balance"
	"github.com/hsson/card-balance-backend/backend/modules/charge"
	"github.com/hsson/card-balance-backend/backend/modules/menu"
	backendConfig "github.com/hsson/card-balance-backend/config"
)

const (
	configFile = "../config/config.yaml"
)

// Run starts the API for appengine
func Run() {
	conf := backendConfig.LoadConfig(configFile)
	menu.Init(conf)
	charge.Init(conf)
	balance.Init(conf)

	r := newRouter()
	http.Handle("/", r)
}
