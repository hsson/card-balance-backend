// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package menu

import (
	"net/http"

	"strings"

	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"
)

// Menu describes a set of restraurants
type Menu struct {
	Language string       `json:"language"`
	Menu     []Restaurant `json:"menu"`
}

// Restaurant describes a restrarant and its dishes
type Restaurant struct {
	Name   string `json:"name"`
	Dishes []Dish `json:"dishes"`
}

// Dish represents a food dish
type Dish struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

// Index gets the entire food menu
func Index(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	lang := vars["lang"]
	urls := getURLs(lang)
	result := Menu{}
	result.Menu = []Restaurant{}
	result.Language = lang
	parser := gofeed.NewParser()
	for _, url := range urls {
		feed, err := parser.ParseURL(url)
		if err != nil {
			return nil, err
		}
		restaurant := Restaurant{}
		restaurant.Name = tidyRestaurantTitle(feed.Title)
		restaurant.Dishes = []Dish{}
		for _, item := range feed.Items {
			dish := Dish{}
			dish.Title = item.Title
			dish.Desc = tidyDishDescription(item.Description)
			restaurant.Dishes = append(restaurant.Dishes, dish)
		}
		result.Menu = append(result.Menu, restaurant)
	}

	return result, nil
}

func tidyDishDescription(menu string) string {
	res := menu
	if strings.Contains(menu, "@") {
		tmp := strings.Split(menu, "@")
		if len(tmp) > 0 {
			res = tmp[0]
		} else {
			res = ""
		}
	}

	if strings.Contains(res, `"`) {
		res = strings.Replace(res, `"`, "", -1)
	}
	return strings.TrimSpace(res)
}

func tidyRestaurantTitle(title string) string {
	return strings.TrimLeft(title, "Meny ")
}
