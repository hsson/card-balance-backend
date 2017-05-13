// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package menu

type restraurantURL struct {
	url      string
	ImageURL string
}

func (u *restraurantURL) getURL(lang string) string {
	return u.url + lang
}

const (
	karenURL      = "http://intern.chalmerskonferens.se/view/restaurant/karrestaurangen/Veckomeny.rss?today=true&locale="
	karenImageURL = ""

	lskitchenURL      = "http://intern.chalmerskonferens.se/view/restaurant/l-s-kitchen/Projektor.rss?today=true&locale="
	lskitchenImageURL = ""

	linsenURL      = "http://intern.chalmerskonferens.se/view/restaurant/linsen/RSS%20Feed.rss?today=true&lang="
	linsenImageURL = ""

	expressURL      = "http://intern.chalmerskonferens.se/view/restaurant/express/V%C3%A4nster.rss?today=true&locale="
	expressImageURL = ""

	hyllanURL      = "http://intern.chalmerskonferens.se/view/restaurant/hyllan/RSS%20Feed.rss?today=true&locale="
	hyllanImageURL = ""
)

var restaurantURLS = []restraurantURL{
	restraurantURL{karenURL, karenImageURL},
	restraurantURL{lskitchenURL, lskitchenImageURL},
	restraurantURL{linsenURL, linsenImageURL},
	restraurantURL{expressURL, expressImageURL},
	restraurantURL{hyllanURL, hyllanImageURL},
}
