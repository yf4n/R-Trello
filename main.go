package main

import (
	"github.com/VojtechVitek/go-trello"
	"github.com/faaaar/R/util"
	"log"
)

func main() {
	util.GetThisWeekDate()

	appKey := util.GetIniConfig("Authorize", "appKey")
	token := util.GetIniConfig("Authorize", "token")
	username := util.GetIniConfig("Authorize", "username")

	trello, err := trello.NewAuthClient(appKey, &token)
	if err != nil {
		log.Fatal(err)
	}

	user, err := trello.Member(username)
	if err != nil {
		log.Fatal(err)
	}

	boards, err := user.Boards()
	if err != nil {
		log.Fatal(err)
	}

	if len(boards) > 0 {
		board := boards[0]
		log.Printf("* %v (%v)\n", board.Name, board.ShortUrl)

		lists, err := board.Lists()
		if err != nil {
			log.Fatal(err)
		}

		for _, list := range lists {
			log.Println("   - ", list.Name)

			cards, _ := list.Cards()
			for _, card := range cards {
				log.Println("      + ", card.Name)
			}
		}
	}
}
