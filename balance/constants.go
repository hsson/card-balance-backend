// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package balance

const (
	// Website info
	baseURL   = "https://kortladdning3.chalmerskonferens.se"
	loginPage = "Default.aspx"
	infoPage  = "CardLoad_Order.aspx"

	// Fields in the login form
	viewStateKey       = "__VIEWSTATE"
	viewStateGenKey    = "__VIEWSTATEGENERATOR"
	eventValidationKey = "__EVENTVALIDATION"
	cardNumberKey      = "txtCardNumber"
	savedNumberKey     = "SavedCardNumber"
	nextButtonKey      = "btnNext"
	nextButtonValue    = "Nästa" // Swedish 'next'
	mobileKey          = "hiddenIsMobile"
	mobileValue        = "desktop"

	// HTTP Headers
	headerContentType        = "Content-Type"
	headerContentTypeValue   = "application/x-www-form-urlencoded"
	headerCookie             = "Cookie"
	headerCookieValue        = "cookieconsent_dismissed=yes"
	headerSetCookie          = "Set-Cookie"
	headerAcceptCharset      = "Accept-Charset"
	headerAcceptCharsetValue = "UTF-8"

	// Regular expression patterns for login form
	viewStatePattern       = "(?:__VIEWSTATE\" value=\")(.*?)(?:\")"
	viewStateGenPattern    = "(?:__VIEWSTATEGENERATOR\" value=\")(.*?)(?:\")"
	eventValidationPattern = "(?:__EVENTVALIDATION\" value=\")(.*?)(?:\")"

	// Regular expression patterns for extracting card data
	cardValuePattern  = "(?:txtPTMCardValue\\\">)(.*?)(?:<\\/span>)"
	cardNamePattern   = "(?:txtPTMCardName\\\">)(.*?)(?:<\\/span>)"
	cardEmailPattern  = "(?:lblEmail\\\">)(.*?)(?:<\\/span>)"
	cardNumberPattern = "(?:txtPTMCardNumber\\\">)(.*?)(?:<\\/span>)"
)
