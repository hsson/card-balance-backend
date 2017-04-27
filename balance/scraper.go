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
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

var (
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
	page, err := s.getWebContent(baseURL + "/" + loginPage)
	if err != nil {
		return Data{}, err
	}
	err = s.updateTokens(page)
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
	if err != nil {
		return Data{}, err
	}
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

func (s *scraper) updateTokens(page string) error {
	htmlNode, err := html.Parse(strings.NewReader(page))
	if err != nil {
		return err
	}
	doc := goquery.NewDocumentFromNode(htmlNode)

	viewState, foundViewState :=
		doc.Find("#" + viewStateKey).First().Attr(valueAttribute)
	viewStateGen, foundViewStateGen :=
		doc.Find("#" + viewStateGenKey).First().Attr(valueAttribute)
	eventValidation, foundEventValidation :=
		doc.Find("#" + eventValidationKey).First().Attr(valueAttribute)

	if !foundViewState || !foundViewStateGen || !foundEventValidation {
		return ErrorNoFormToken
	}

	s.viewState = viewState
	s.viewStateGen = viewStateGen
	s.eventValidation = eventValidation
	return nil
}

func parseData(page string) (Data, error) {
	htmlNode, err := html.Parse(strings.NewReader(page))
	if err != nil {
		return Data{}, err
	}
	doc := goquery.NewDocumentFromNode(htmlNode)
	data := Data{}
	data.CardNumber = doc.Find(cardNumberID).First().Text()
	data.FullName = doc.Find(cardNameID).First().Text()
	data.Email = doc.Find(cardEmailID).First().Text()
	balance := doc.Find(cardValueID).First().Text()
	value, err := strconv.ParseFloat(strings.Replace(balance, ",", ".", -1), 64)
	if err != nil {
		return Data{}, ErrorInvalidBalance
	}
	data.Balance = value
	return data, nil
}
