package main

import (
	"fmt"
	"github.com/VojtechVitek/go-trello"
	Util "github.com/faaaar/R/util"
	"log"
)

func main() {
	appKey := Util.GetIniConfig("Authorize", "appKey")
	token := Util.GetIniConfig("Authorize", "token")
	username := Util.GetIniConfig("Authorize", "username")

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
		fmt.Printf("* %v (%v)\n", board.Name, board.ShortUrl)

		lists, err := board.Lists()
		if err != nil {
			log.Fatal(err)
		}

		for _, list := range lists {
			fmt.Println("   - ", list.Name)

			cards, _ := list.Cards()
			for _, card := range cards {
				fmt.Println("      + ", card.Name)
			}
		}
	}
}
