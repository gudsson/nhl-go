package scrapers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gudsson/nhl-go/model"
	"github.com/gudsson/nhl-go/utils"
)

func GetEvents(gameId string) []model.Event {
	year, _ := strconv.Atoi(gameId[0:4])
	pbpId := []string{strconv.Itoa(year) + strconv.Itoa(year + 1), os.Args[1][4:]}
	fmt.Println(pbpId)

	c := colly.NewCollector(
		colly.AllowedDomains("www.nhl.com"),
	)

	eventArr := []model.Event{}

	c.OnHTML("tr[id^='PL-']", func(e *colly.HTMLElement) {
		tds := []*goquery.Selection{}

		e.DOM.Children().Each(func(i int, s *goquery.Selection) {
			tds = append(tds, s)
		})

		
		eventId, _ := strconv.Atoi(tds[0].Text())
		periodNum, _ := strconv.Atoi(tds[1].Text())
		periodTimes := utils.PeriodClock(tds[3].Text())

		awayPlayers := utils.PlayersOnIce(tds[6])
		homePlayers := utils.PlayersOnIce(tds[7])

		event := model.Event{
			EventId: int(eventId),
			Period: int(periodNum),
			Strength: tds[2].Text(),
			PeriodTimeElapsed: periodTimes[0],
			PeriodTimeRemaining: periodTimes[1],
			EventCode: tds[4].Text(),
			EventDescription: tds[5].Text(),
			AwayOnIce: awayPlayers,
			HomeOnIce: homePlayers,
		}
		
		eventArr = append(eventArr, event)
	})

	url := fmt.Sprintf("http://www.nhl.com/scores/htmlreports/%s/PL%s.HTM", pbpId[0], pbpId[1])
	c.Visit(url)

	return eventArr
}