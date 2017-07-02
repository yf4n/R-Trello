package main

import (
	"github.com/VojtechVitek/go-trello"
	"github.com/faaaar/R/util"
	"log"
	"time"
)

var p = log.Println

func main() {
	appKey := util.GetIniConfig("authorize", "appKey")
	token := util.GetIniConfig("authorize", "token")
	username := util.GetIniConfig("filter", "username")
	outputPath := util.GetIniConfig("path", "output")

	trello, err := trello.NewAuthClient(appKey, &token)
	util.CheckError(err)

	user, _ := trello.Member(username)
	userid := user.Id

	boards, _ := user.Boards()

	timeNow := time.Now()
	weekTimestampRange := util.GetWeekDateRange(timeNow)
	startTs := weekTimestampRange["startTs"]
	endTs := weekTimestampRange["endTs"]
	startTime := util.GetDateStringWithFormat(startTs, "2006-01-02")
	endTime := util.GetDateStringWithFormat(endTs, "2006-01-02")
	markdownStr := "# " + startTime + " - " + endTime + " 工作内容 \n"

	for _, board := range boards {
		// boardName := board.Name
		cards, _ := board.Cards()

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
					t, _ := time.Parse(time.RFC3339, card.Due)
					ts := t.Unix()

					if ts > startTs && ts < endTs {
						markdownStr = markdownStr + "## " + card.Name + "（完成时间: " + util.GetDateString(ts) + ")" + "\n"
						markdownStr = markdownStr + card.Desc + "\n"

						lists, _ := card.Checklists()
						for _, list := range lists {
							// p(list.Name)
							markdownStr += "### " + list.Name + "\n"

							for _, item := range list.CheckItems {
								// p(item.Name)
								markdownStr += "- " + item.Name + "\n"
							}

							markdownStr += "\n"
						}
					}
				}
			}
		}
	}
	markdownStr += "--\n"
	markdownStr += "*此周报由 周报生成器0.1 生成*\n"
	markdownStr += "*开源地址: https://github.com/faaaar/R*\n"

	util.WriteFile(outputPath+"/"+util.GetTodayDateString()+".md", markdownStr)
	util.SendMail(outputPath+"/"+util.GetTodayDateString()+".md", "["+startTime+"-"+endTime+"]")
}
