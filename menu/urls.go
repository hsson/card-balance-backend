// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package menu

var urls = []string{"http://intern.chalmerskonferens.se/view/restaurant/karrestaurangen/Veckomeny.rss?today=true&locale=", "http://intern.chalmerskonferens.se/view/restaurant/l-s-kitchen/Projektor.rss?today=true&locale=", "http://intern.chalmerskonferens.se/view/restaurant/linsen/RSS%20Feed.rss?today=true&lang=", "http://intern.chalmerskonferens.se/view/restaurant/express/V%C3%A4nster.rss?today=true&locale=", "http://intern.chalmerskonferens.se/view/restaurant/hyllan/RSS%20Feed.rss?today=true&locale="}

func getURLs(lang string) []string {
	results := []string{}
	for _, url := range urls {
		results = append(results, url+lang)
	}
	return results
}
