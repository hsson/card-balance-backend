// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package balance

import (
	"errors"
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

type scraper struct {
	cardNumber string

	viewState       string
	viewStateGen    string
	eventValidation string

	headers map[string]string

	client *http.Client
}

func (s *scraper) init() {
	s.cardNumber = ""
	s.viewState = ""
	s.viewStateGen = ""
	s.eventValidation = ""
	s.headers = nil
	s.client = &http.Client{
		Timeout: time.Second * 10,
	}
}

func (s *scraper) Scrape(number string) (Data, error) {
	// Initialize scraper for new scrape request
	s.init()
	s.cardNumber = number

	// Get tokens from login form
	err := s.updateTokens()
	if err != nil {
		return Data{}, err
	}

	// Perform the login
	err = s.login()
	if err != nil {
		return Data{}, err
	}

	// Get the data from the logged in page
	data, err := s.getData()
	return data, nil
}

func (s *scraper) login() error {
	// Prepare a request with correct headers and login
	// form values
	req, err := s.newLoginRequest()
	if err != nil {
		return err
	}
	// Make sure to not follow any redirects
	s.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	response, err := s.client.Do(req)
	if err != nil {
		return err
	}
	// Extract the session cookie from the response
	cookie := response.Header.Get(headerSetCookie)
	if cookie == "" {
		return ErrorNoSessionCookie
	}
	s.headers = make(map[string]string)
	s.headers[headerCookie] = cookie
	s.headers[headerAcceptCharset] = headerAcceptCharsetValue
	return nil
}

func (s *scraper) getData() (Data, error) {
	page, err := s.getWebContent(baseURL + "/" + infoPage)
	if err != nil {
		return Data{}, err
	}
	data, err := parseData(page)
	if err != nil {
		return Data{}, err
	}
	return data, nil
}

func (s *scraper) getWebContent(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	if s.headers != nil {
		for key, value := range s.headers {
			req.Header.Add(key, value)
		}
	}
	response, err := s.client.Do(req)
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

func (s *scraper) newLoginRequest() (*http.Request, error) {
	// Add login form parameters
	loginForm := url.Values{}
	loginForm.Add(viewStateKey, s.viewState)
	loginForm.Add(viewStateGenKey, s.viewStateGen)
	loginForm.Add(eventValidationKey, s.eventValidation)
	loginForm.Add(cardNumberKey, s.cardNumber)
	loginForm.Add(savedNumberKey, "")
	loginForm.Add(nextButtonKey, nextButtonValue)
	loginForm.Add(mobileKey, mobileValue)

	// Create an HTTP request with the login parameters
	req, err := http.NewRequest(http.MethodPost, baseURL+"/"+loginPage, strings.NewReader(loginForm.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add(headerContentType, headerContentTypeValue)
	req.Header.Add(headerCookie, headerCookieValue)
	return req, nil
}

func (s *scraper) updateTokens() error {
	page, err := s.getWebContent(baseURL + "/" + loginPage)
	if err != nil {
		return err
	}

	s.viewState = extractData(page, viewStateRegexp)
	s.viewStateGen = extractData(page, viewStateGenRegexp)
	s.eventValidation = extractData(page, eventValidationRegexp)
	if s.viewState == "" || s.viewStateGen == "" || s.eventValidation == "" {
		return ErrorNoFormToken
	}
	return nil
}

func parseData(page string) (Data, error) {
	data := Data{}
	// TODO: CHeck so values are not empty
	valText := extractData(page, cardValueRegexp)
	name := extractData(page, cardNameRegexp)
	email := extractData(page, cardEmailRegexp)
	cardNumber := extractData(page, cardNumberRegexp)
	value, err := strconv.ParseFloat(strings.Replace(valText, ",", ".", -1), 64)
	if err != nil {
		return Data{}, ErrorInvalidBalance
	}

	data.FullName = name
	data.Email = email
	data.CardNumber = cardNumber
	data.Balance = value
	return data, nil
}

func extractData(input string, expression *regexp.Regexp) string {
	results := expression.FindStringSubmatch(input)
	if len(results) <= 1 {
		return ""
	}
	return results[1]
}
