// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package backend

import (
	"github.com/hsson/card-balance-backend/backend/modules/balance"
	"github.com/hsson/card-balance-backend/backend/modules/charge"
	"github.com/hsson/card-balance-backend/backend/modules/menu"
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
	// Balance V2
	route{"GetBalanceV2", "GET", "/balance/v2/{number:[0-9]+}/{userInfo}", balance.GetBalanceV2},
	// Food menu module
	route{"IndexMenu", "GET", "/menu/{lang:(?:sv|en)}", menu.Index},
	// Card charging redirect
	route{"ChargeRedirect", "GET", "/charge", charge.Redirect},
}
