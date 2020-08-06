package main

import (
	"github.com/gocolly/colly/v2"
)

func lebonScrap(url string, cout chan cardStruct) {

	c := colly.NewCollector()
	card := cardStruct{URL: url}
	defer c.Visit(card.URL)

	c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36"

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		card.Title = e.Text
	})

	c.OnHTML(`div[data-qa-id='adview_price']`, func(e *colly.HTMLElement) {
		card.Price = e.Text
	})

	c.OnHTML("img[alt='image-galerie-0']", func(e *colly.HTMLElement) {
		card.Picture = e.Attr("src")
	})

	c.OnScraped(func(r *colly.Response) {
		cout <- card
	})
}
