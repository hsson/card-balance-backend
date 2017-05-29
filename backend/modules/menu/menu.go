// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package menu

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hsson/card-balance-backend/backend/modules"
	backendConfig "github.com/hsson/card-balance-backend/config"
)

// Menu describes a set of restraurants
type Menu struct {
	Language string       `json:"language"`
	Menu     []Restaurant `json:"menu"`
}

// Restaurant describes a restrarant and its dishes
type Restaurant struct {
	Name         string     `json:"name"`
	ImageURL     string     `json:"image_url"`
	WebsiteURL   string     `json:"website_url"`
	Rating       float32    `json:"rating"`
	AveragePrice int        `json:"avg_price"`
	Campus       string     `json:"campus"`
	Dishes       []Dish     `json:"dishes"`
	OpenHours    []OpenHour `json:"open_hours"`
}

// Dish represents a food dish
type Dish struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type jsonRestaurant struct {
	MenuDate         string `json:"menuDate"`
	RecipeCategories []struct {
		Name        string `json:"name"`
		NameEnglish string `json:"nameEnglish"`
		ID          int    `json:"id"`
		Recipes     []struct {
			DisplayNames []struct {
				TypeID      int    `json:"typeID"`
				DisplayName string `json:"displayName"`
			} `json:"displayNames"`
			CO2E      string `json:"cO2e"`
			CO2EURL   string `json:"cO2eURL"`
			Allergens []struct {
				ID             int         `json:"id"`
				ImageURLBright interface{} `json:"imageURLBright"`
				ImageURLDark   interface{} `json:"imageURLDark"`
			} `json:"allergens"`
			Price float64 `json:"price"`
		} `json:"recipes"`
	} `json:"recipeCategories"`
}

const (
	languageSwedish = "sv"
)

// OpenHour represents the open hour of a restaurant
type OpenHour struct {
	DayOfWeek int `json:"day_of_week"`
	StartTime int `json:"start_hour"`
	EndTime   int `json:"end_hour"`
}

var config backendConfig.Config

// Init initializes the module with specified config
func Init(newConfig backendConfig.Config) {
	config = newConfig
}

// Index gets the entire food menu using the JSON endpoint
func Index(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	lang := vars["lang"]
	result := Menu{}
	result.Menu = []Restaurant{}
	result.Language = lang

	httpClient := modules.GetHTTPClient(r)
	for _, rawRestaurant := range config.Restaurants {
		data, err := httpClient.Get(rawRestaurant.MenuURL)
		if err != nil {
			return nil, err
		}
		decoder := json.NewDecoder(data.Body)
		jsonResponse := jsonRestaurant{}
		err = decoder.Decode(&jsonResponse)
		if err != nil {
			return nil, err
		}
		restaurant := Restaurant{}
		restaurant.Name = rawRestaurant.Name
		restaurant.Dishes = []Dish{}
		restaurant.ImageURL = rawRestaurant.ImageURL
		restaurant.WebsiteURL = rawRestaurant.WebsiteURL
		restaurant.Rating = rawRestaurant.Rating
		for _, category := range jsonResponse.RecipeCategories {
			for _, recipe := range category.Recipes {
				dish := Dish{}
				if lang == languageSwedish {
					dish.Title = category.Name
				} else {
					dish.Title = category.NameEnglish
				}
				if len(recipe.DisplayNames) == 1 {
					dish.Desc = recipe.DisplayNames[0].DisplayName
				} else {
					if lang == languageSwedish {
						dish.Desc = recipe.DisplayNames[0].DisplayName
					} else {
						dish.Desc = recipe.DisplayNames[1].DisplayName
					}
				}
				restaurant.Dishes = append(restaurant.Dishes, dish)
			}
		}
		restaurant.AveragePrice = rawRestaurant.AveragePrice
		restaurant.Campus = rawRestaurant.Campus
		openHours := []OpenHour{}
		for _, oh := range rawRestaurant.OpenHours {
			openHour := OpenHour{}
			openHour.DayOfWeek = oh.DayOfWeek
			openHour.StartTime = oh.StartTime
			openHour.EndTime = oh.EndTime
			openHours = append(openHours, openHour)
		}
		restaurant.OpenHours = openHours
		result.Menu = append(result.Menu, restaurant)
	}

	return result, nil
}
