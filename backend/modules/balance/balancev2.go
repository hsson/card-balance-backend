package balance

import (
	"errors"
	"net/http"

	"github.com/hsson/card-balance-backend/backend/modules"

	"github.com/gorilla/mux"
)

// ErrorInvalidUserInfo is sent when the client didn't specify any userinfo
var ErrorInvalidUserInfo = errors.New("No user info set")

// GetBalanceV2 is a new version of the balance API that depends on some
// userInfo cookie being set in the client
func GetBalanceV2(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	cardNumber := vars["number"]
	if len(cardNumber) != allowedCardNumberLength {
		return nil, ErrorInvalidCardNumber
	}
	userInfo := vars["userInfo"]
	if len(userInfo) == 0 {
		return nil, ErrorInvalidUserInfo
	}

	scraper := new(scraper)
	scraper.init()
	scraper.client = modules.GetHTTPClient(r)
	data, err := scraper.Scrape(cardNumber, userInfo)
	if err != nil {
		return nil, err
	}
	return data, nil
}
