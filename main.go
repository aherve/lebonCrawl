package main

import (
	"log"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please provide an url as first argument")
	}
	url := os.Args[1]

	reg := regexp.MustCompile(`www\.leboncoin\.fr`)
	if reg.MatchString(url) {
		cout := make(chan cardStruct)
		go lebonScrap(url, cout)
		card := <-cout

		if err := card.Export(); err != nil {
			log.Fatal(err)
		}
		log.Println("Card successfully created")
	} else {
		log.Fatal("unrecognized url")
	}
}
