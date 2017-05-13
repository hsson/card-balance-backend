// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package backend

import (
	"github.com/hsson/card-balance-backend/balance"
	"github.com/hsson/card-balance-backend/charge"
	"github.com/hsson/card-balance-backend/menu"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc errorHandler
}

var routes = []route{
	// Balance module
	route{"GetBalance", "GET", "/balance/{number:[0-9]+}", balance.GetBalance},
	// Food menu module
	route{"IndexMenu", "GET", "/menu/{lang:(sv|en)}", menu.Index},
	// Card charging redirect
	route{"ChargeRedirect", "GET", "/charge", charge.Redirect},
}
