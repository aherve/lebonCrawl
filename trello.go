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
	//boardID := os.Getenv("TRELLO_BOARD_ID")
	listID := os.Getenv("TRELLO_LIST_ID")

	client := trello.NewClient(key, token)

	// Create card (without attachment, unfortunately)
	err := client.CreateCard(&trello.Card{
		Name:   card.Title,
		Desc:   card.URL,
		IDList: listID,
	}, trello.Defaults())
	if err != nil {
		return err
	}
	log.Println("Card created. retrieving id to add attachment...")

	// Retrieve card and add attachment. Tedious: create should return an ID !
	list, err := client.GetList(listID, trello.Defaults())
	if err != nil {
		return err
	}
	cards, err := list.GetCards(trello.Defaults())
	if err != nil {
		return err
	}
	for _, c := range cards {
		if c.Desc == card.URL {
			log.Println("Card found. adding attachment...")
			return c.AddURLAttachment(&trello.Attachment{
				URL: card.Picture,
			})
		}
	}

	return nil

}
