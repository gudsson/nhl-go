package scrapers

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gudsson/nhl-go/model"
	"github.com/gudsson/nhl-go/utils"
)

func GetEvents() []model.Event {
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

	c.Visit("http://www.nhl.com/scores/htmlreports/20212022/PL020562.HTM") 

	return eventArr
}