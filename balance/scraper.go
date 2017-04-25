// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package balance

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	// Regex foŕ the login form
	viewStateRegexp       = regexp.MustCompile(viewStatePattern)
	viewStateGenRegexp    = regexp.MustCompile(viewStateGenPattern)
	eventValidationRegexp = regexp.MustCompile(eventValidationPattern)

	// Regex for extracting card data
	cardValueRegexp  = regexp.MustCompile(cardValuePattern)
	cardNameRegexp   = regexp.MustCompile(cardNamePattern)
	cardEmailRegexp  = regexp.MustCompile(cardEmailPattern)
	cardNumberRegexp = regexp.MustCompile(cardNumberPattern)

	// ErrorBadPage indicates that the page is not available
	ErrorBadPage = errors.New("Page not OK")
	// ErrorNoFormToken means that a required form token is not available
	ErrorNoFormToken = errors.New("A required form token is unavailable")
	// ErrorNoSessionCookie means a session cookie could not be obtained
	ErrorNoSessionCookie = errors.New("Session cookie could not be retrieved")
	// ErrorInvalidBalance means that the balance could not be parsed
	ErrorInvalidBalance = errors.New("The balance could not be retrieved")
)

type formTokens struct {
	viewState          string
	viewStateGenerator string
	eventValidation    string
}

func getCardData(number string) (Data, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	// Request to get the login page, this is needed as
	// some tokens are required from the login form
	loginPage, err := getPage(netClient, baseURL+"/"+loginPage, nil)
	if err != nil {
		return Data{}, err
	}
	// Extract the tokens form the login form
	tokens, err := getFormTokens(loginPage)
	if err != nil {
		return Data{}, err
	}
	// Do the login and get a session cookie
	cookie, err := getSessionCookie(netClient, number, tokens)
	if err != nil {
		return Data{}, err
	}
	headers := make(map[string]string)
	headers[headerCookie] = cookie
	headers[headerAcceptCharset] = headerAcceptCharsetValue
	detailsPage, err := getPage(netClient, baseURL+"/"+infoPage, headers)
	if err != nil {
		return Data{}, err
	}

	// TODO: Check so values aren't empty
	valText := extractData(detailsPage, cardValueRegexp)
	name := extractData(detailsPage, cardNameRegexp)
	email := extractData(detailsPage, cardEmailRegexp)
	cardNumber := extractData(detailsPage, cardNumberRegexp)
	value, err := strconv.ParseFloat(strings.Replace(valText, ",", ".", -1), 64)
	if err != nil {
		return Data{}, ErrorInvalidBalance
	}

	data := Data{}
	data.CardNumber = cardNumber
	data.FullName = name
	data.Email = email
	data.Balance = value
	fmt.Println(value)
	return data, nil
}

func getPage(client *http.Client,
	url string,
	headers map[string]string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		return "", ErrorBadPage
	}
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func getFormTokens(loginPage string) (formTokens, error) {
	tokens := formTokens{}
	tokens.viewState = extractData(loginPage, viewStateRegexp)
	tokens.viewStateGenerator = extractData(loginPage, viewStateGenRegexp)
	tokens.eventValidation = extractData(loginPage, eventValidationRegexp)
	if tokens.viewState == "" || tokens.viewStateGenerator == "" || tokens.eventValidation == "" {
		return tokens, ErrorNoFormToken
	}
	return tokens, nil
}

func getSessionCookie(client *http.Client,
	cardNumber string,
	tokens formTokens) (string, error) {

	formValues := url.Values{}
	// Add form parameters
	formValues.Add(viewStateKey, tokens.viewState)
	formValues.Add(viewStateGenKey, tokens.viewStateGenerator)
	formValues.Add(eventValidationKey, tokens.eventValidation)
	formValues.Add(cardNumberKey, cardNumber)
	formValues.Add(savedNumberKey, "")
	formValues.Add(nextButtonKey, nextButtonValue)
	formValues.Add(mobileKey, mobileValue)

	req, err := http.NewRequest(http.MethodPost, baseURL+"/"+loginPage, strings.NewReader(formValues.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Add(headerContentType, headerContentTypeValue)
	req.Header.Add(headerCookie, headerCookieValue)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	cookie := response.Header.Get(headerSetCookie)
	if cookie == "" {
		return "", ErrorNoSessionCookie
	}
	return cookie, nil
}

func extractData(input string, expression *regexp.Regexp) string {
	results := expression.FindStringSubmatch(input)
	if len(results) <= 1 {
		return ""
	}
	return results[1]
}
