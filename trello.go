package main

import (
	"github.com/adlio/trello"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type cardStruct struct {
	Price   string
	URL     string
	Title   string
	Picture string
}

func init() {
	godotenv.Load("./secret.env")
}

func (card cardStruct) Export() error {
	log.Println("exporting trello card")
	token := os.Getenv("TRELLO_TOKEN")
	key := os.Getenv("TRELLO_KEY")
	listID := os.Getenv("TRELLO_LIST_ID")

	client := trello.NewClient(key, token)

	// Create card (without attachment, unfortunately)
	tCard := trello.Card{
		Name:   card.Title,
		Desc:   card.URL,
		IDList: listID,
	}
	err := client.CreateCard(&tCard, trello.Defaults())
	if err != nil {
		return err
	}
	log.Println("Card created. adding attachment...")
	return tCard.AddURLAttachment(&trello.Attachment{
		URL: card.Picture,
	})
}
