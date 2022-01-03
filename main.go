package main

// PBP: http://www.nhl.com/scores/htmlreports/20212022/PL020562.HTM
// live: https://statsapi.web.nhl.com/api/v1/game/2021020562/feed/live

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gudsson/nhl-go/model"
	"github.com/gudsson/nhl-go/utils"
)

// type people struct {
// 	Number int `json:"number"`
// }

func main() {
	// events := getEvents()
	// text := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft": "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft": "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`

	// url := "http://api.open-notify.org/astros.json"

	// textBytes := []byte(text)

	// people1 := people{}
	// err := json.Unmarshal(textBytes, &people1)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(people1.Number)
	// events := getEvents()
	// fmt.Println(events)
	getCoordinates()
}

func getCoordinates() {
	// c := colly.NewCollector(
	// 	colly.AllowedDomains("statsapi.web.nhl.com"),
	// )

	// c.OnHTML("body", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.DOM.Text())
	// })

	// c.Visit("https://statsapi.web.nhl.com/api/v1/game/2021020562/feed/live") 
	resp, err := http.Get("https://statsapi.web.nhl.com/api/v1/game/2021020562/feed/live")
	if err != nil {
			fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result model.LiveFeedJSON
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	// fmt.Println(string(body))              // convert to string before print
	fmt.Println(result.Copyright)
}

func getEvents() []model.Event {
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
	})

	c.Visit("http://www.nhl.com/scores/htmlreports/20212022/PL020562.HTM") 

	return eventArr
}