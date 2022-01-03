package main

// PBP: http://www.nhl.com/scores/htmlreports/20212022/PL020562.HTM
// live: https://statsapi.web.nhl.com/api/v1/game/2021020562/feed/live

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gudsson/nhl-go/model"
	"github.com/gudsson/nhl-go/utils"
)

func main() {


	events := getData()

	fmt.Println(events[6])
}

func getData() []model.Event {
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
			EventId: int16(eventId),
			Period: int8(periodNum),
			Strength: tds[2].Text(),
			PeriodTimeElapsed: periodTimes[0],
			PeriodTimeRemaining: periodTimes[1],
			EventCode: tds[4].Text(),
			EventDescription: tds[5].Text(),
			AwayOnIce: awayPlayers,
			HomeOnIce: homePlayers,
		}
		
		eventArr = append(eventArr, event)
		// fmt.Println(events)
	})

	// fmt.Println(events)

	c.Visit("http://www.nhl.com/scores/htmlreports/20212022/PL020562.HTM") 

	return eventArr
}

// func awaitTask() <-chan int32 {
// 	r := make(chan int32)

// 	go func() {
// 		defer close(r)


// 	}

// 	c := colly.NewCollector(
// 			colly.AllowedDomains("www.nhl.com"),
// 	)

// 	events := []model.Event{}

// 	c.OnHTML("tr[id^='PL-']", func(e *colly.HTMLElement) {
// 		tds := []*goquery.Selection{}

// 		e.DOM.Children().Each(func(i int, s *goquery.Selection) {
// 			tds = append(tds, s)
// 		})

		
// 		eventId, _ := strconv.Atoi(tds[0].Text())
// 		periodNum, _ := strconv.Atoi(tds[1].Text())
// 		periodTimes := utils.PeriodClock(tds[3].Text())

// 		awayPlayers := utils.PlayersOnIce(tds[6])
// 		homePlayers := utils.PlayersOnIce(tds[7])

// 		event := model.Event{
// 			EventId: int16(eventId),
// 			Period: int8(periodNum),
// 			Strength: tds[2].Text(),
// 			PeriodTimeElapsed: periodTimes[0],
// 			PeriodTimeRemaining: periodTimes[1],
// 			EventCode: tds[4].Text(),
// 			EventDescription: tds[5].Text(),
// 			AwayOnIce: awayPlayers,
// 			HomeOnIce: homePlayers,
// 		}
		
// 		events = append(events, event)
// 	})

// 	// fmt.Println(events)

// 	c.Visit("http://www.nhl.com/scores/htmlreports/20212022/PL020562.HTM") 
// }