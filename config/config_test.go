package backend

import (
	"io/ioutil"
	"testing"

	"net/url"

	"strings"

	"gopkg.in/yaml.v2"
)

var dummyConfig = `
# Configuration file for the backend
card_site_url: https://example.com/cardstuff

restaurants:
  - name: Kårrestaurangen
    menu_url: https://example.com/kårrestaurangen
    image:
    website:
    rating: 3.0
    open_hours:
      - day_of_week: 2
        start_hour: 800
        end_hour: 1330

  - name: L's Kitchen
    menu_url: https://example.com/lskitchen
    image: http://example.com/some/image
    website: https://example.com/thewebsite
    rating: 1.0
`

func TestLoadDummyConfig(t *testing.T) {
	config := Config{}
	err := yaml.Unmarshal([]byte(dummyConfig), &config)
	if err != nil {
		t.Error("Could not parse config")
	}

	if config.CardSiteURL != "https://example.com/cardstuff" {
		t.Error("CardSiteURL is incorrect")
	}

	if size := len(config.Restaurants); size != 2 {
		t.Errorf("There should be 2 restaurants, there are %d\n", size)
	}

	if config.Restaurants[0].Name != "Kårrestaurangen" {
		t.Error("Incorrect name of first restaurant")
	}

	if config.Restaurants[1].Rating != 1.0 {
		t.Error("Incorrect rating of second restaurant")
	}
}

func TestLoadRealConfig(t *testing.T) {
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		t.Error("Could not load config file")
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		t.Error("Error when parsing config")
	}

	if !strings.HasSuffix(config.CardSiteURL, "/") {
		t.Error("The card site URL should end with a '/'")
	}

	// Test that all URLs are valid. Image URL and website URL
	// are optional
	for _, res := range config.Restaurants {
		_, urlErr := url.ParseRequestURI(res.MenuURL)
		if urlErr != nil {
			t.Errorf("Not a valid url: %v\n", res.MenuURL)
		}
		_, urlErr = url.ParseRequestURI(res.ImageURL)
		if res.ImageURL != "" && urlErr != nil {
			t.Errorf("Not a valid url: %v\n", res.ImageURL)
		}
		_, urlErr = url.ParseRequestURI(res.WebsiteURL)
		if res.WebsiteURL != "" && urlErr != nil {
			t.Errorf("Not a valid url: %v\n", res.WebsiteURL)
		}

		// Test rating range
		if res.Rating < 0.0 || res.Rating > 5.0 {
			t.Errorf("Rating must be between 0 and 5, is: %v\n", res.Rating)
		}
	}
}
