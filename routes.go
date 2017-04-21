// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"net/http"

	"github.com/hsson/card-balance-backend/balance"
	"github.com/hsson/card-balance-backend/charge"
	"github.com/hsson/card-balance-backend/menu"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []route{
	// Balance module
	route{"GetBalance", "GET", "/balance/{number:[0-9]{16}}", balance.GetBalance},
	// Food menu module
	route{"IndexMenu", "GET", "/menu", menu.Index},
	// Card charging redirect
	route{"ChargeRedirect", "GET", "/charge", charge.Redirect},
}
