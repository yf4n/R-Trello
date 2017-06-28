package main

import (
	"github.com/VojtechVitek/go-trello"
	"github.com/faaaar/R/util"
	"log"
	"time"
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
	userid := user.Id

	boards, err := user.Boards()
	if err != nil {
		log.Fatal(err)
	}

	startDay := time.Now()

	for startDay.Weekday() != time.Monday {
		startDay = startDay.AddDate(0, 0, -1)
	}

	startStr := startDay.Format("2006-01-02")
	startF, _ := time.Parse("2006-01-02", startStr)
	startTs := startF.Unix()

	log.Println(startTs)
	for _, board := range boards {
		boardName := board.Name
		log.Printf("* %v (%v)\n", boardName, board.ShortUrl)

		cards, err := board.Cards()
		if err != nil {
			log.Fatal(err)
		}

		for _, card := range cards {
			isOwnCard := false
			memberIdList := card.IdMembers
			for _, id := range memberIdList {
				if id == userid {
					isOwnCard = true
					break
				}
			}

			if isOwnCard {
				if card.Due != "" {
					t, err := time.Parse(time.RFC3339, card.Due)
					if err != nil {
						log.Fatalln(err)
					}

					log.Println(t.Weekday(), time.Now().Unix())
				}
			}
		}
		// for _, list := range lists {
		//   log.Println("   - ", list.Name)
		//
		//   cards, _ := list.Cards()
		//   for _, card := range cards {
		//     log.Println("      + ", card.Name)
		//   }
		// }
	}
}
