package main

import (
	"log"
	"time"

	"github.com/VojtechVitek/go-trello"
	"github.com/faaaar/R/util"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var p = log.Println
var appKey = util.GetIniConfig("authorize", "appKey")
var token = util.GetIniConfig("authorize", "token")
var username = util.GetIniConfig("filter", "username")
var filterBoard = util.GetIniConfig("filter", "board")
var outputPath = util.GetIniConfig("path", "output")

func main() {
	p("请求&&处理数据中...")
	trelloClient, err := trello.NewAuthClient(appKey, &token)
	util.CheckError(err)

	markdownStr := ""
	markdownStr += generateReportTitle()
	markdownStr += generateCurrentWeekReport(trelloClient)
	markdownStr += generateNextWeekReport(trelloClient)
	markdownStr += "--\n"
	markdownStr += "*此周报由 周报生成器 0.2 生成*\n"
	markdownStr += "*开源地址: https://github.com/faaaar/R*\n"

	p(markdownStr)
	p("正在生成markdown到" + outputPath + "...")
	util.WriteFile(outputPath+"/"+util.GetTodayDateString()+".md", markdownStr)
	unsafe := blackfriday.MarkdownCommon([]byte(markdownStr))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	util.WriteFile(outputPath+"/"+util.GetTodayDateString()+".html", string(html))

	// p("正在发送邮件...")
	// util.SendMail(outputPath+"/"+util.GetTodayDateString()+".md", "("+startTime+" - "+endTime+")")

	p("完成...")
}

func generateReportTitle() string {
	timeNow := time.Now()
	startTs, endTs := util.GetWeekDateRange(timeNow)
	startTime := util.GetDateStringWithFormat(startTs, "2006/01/02")
	endTime := util.GetDateStringWithFormat(endTs, "2006/01/02")

	return "## " + startTime + " - " + endTime + " 周报 \n\n"
}

// generateCurrentWeekReport 生成本周工作内容
func generateCurrentWeekReport(trelloClient *trello.Client) string {
	startTs, endTs := util.GetWeekDateRange(time.Now())

	return "### 本周工作内容 \n\n" + generateWeekReport(trelloClient, startTs, endTs)
}

// generateNextWeekReport 生成下周工作计划
func generateNextWeekReport(trelloClient *trello.Client) string {
	startTs, endTs := util.GetWeekDateRange(time.Now().AddDate(0, 0, 7))

	return "### 下周工作计划 \n\n" + generateWeekReport(trelloClient, startTs, endTs)
}

// generateWeekReport 生成startTs和endTs时间范围内的工作内容
func generateWeekReport(trello *trello.Client, startTs int64, endTs int64) string {
	user, _ := trello.Member(username)
	userid := user.Id

	boards, _ := user.Boards()

	markdownStr := ""

	for _, board := range boards {
		boardName := board.Name

		if boardName != filterBoard {
			continue
		}

		cards, _ := board.Cards()

		for _, card := range cards {
			isOwnCard := false
			memberIDList := card.IdMembers
			for _, id := range memberIDList {
				if id == userid {
					isOwnCard = true
					break
				}
			}

			if isOwnCard {
				if card.Due != "" {
					t, _ := time.Parse(time.RFC3339, card.Due)
					ts := t.Unix()
					if ts > startTs && ts <= endTs {
						typeStr := ""
						if len(card.Labels) > 0 {
							typeStr += "| "
						}

						for _, label := range card.Labels {
							typeStr += label.Name + " | "
						}
						markdownStr = markdownStr + "### " + card.Name + "（完成时间: " + util.GetDateStringWithFormat(ts, "2006/01/02") + ")" + "\n\n"
						markdownStr += typeStr + "\n\n"
						markdownStr += card.Desc + "\n\n"

						lists, _ := card.Checklists()
						for _, list := range lists {
							markdownStr += "#### " + list.Name + "\n"

							for _, item := range list.CheckItems {
								markdownStr += "- " + item.Name + "\n"
							}

							markdownStr += "\n"
						}
					}
				}
			}
		}
	}

	return markdownStr
}
