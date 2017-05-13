// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build !appengine

package main

import (
	"github.com/hsson/card-balance-backend/backend"
)

func main() {
	backend.Run()
}
