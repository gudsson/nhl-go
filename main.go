package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gudsson/nhl-go/utils"
)

func main() {
	c := colly.NewCollector(
			colly.AllowedDomains("www.nhl.com"),
	)

	// in table row for event
	// - 8 first level tds
	//  - idx 0: event number
	//  - idx 1: period unmber
	//  - idx 2: strength state
	//  - idx 3: time elapsed <br /> remaining
	//  - idx 4: event code
	//  - idx 5: event description
	//  - idx 6: away on ice
	//    - td
	//      - table
	//        - tbody
	//          - tr
	//            - td
	//              - table
	//                - tbody
	//                  - tr
	//                    - td > PLAYER_NUMBER
	//                      - font [style='cursor:hand'] > PLAYER_POS PLAYER_NAME
  //                  - tr
	//                    - td > PLAYER_POS
  //  - idx 7: home on ice
	//    - td
	//      - table
	//        - tbody
	//          - tr
	//            - td
	//              - table
	//                - tbody
	//                  - tr
	//                    - td > PLAYER_NUMBER
	//                      - font [style='cursor:hand'] > PLAYER_POS PLAYER_NAME
  //                  - tr
	//                    - td > PLAYER_POS  

	c.OnHTML("tr[id^='PL-']", func(e *colly.HTMLElement) {
		tds := []*goquery.Selection{}

		e.DOM.Children().Each(func(i int, s *goquery.Selection) {
			tds = append(tds, s)
		})

		for i, v := range tds {
			if i == 6 { // away team
				v.Find("font").Each(func(_ int, s *goquery.Selection) {
					playerStr, _ := s.Attr("title")
					playerSlice := utils.PlayerPosAndName(playerStr)
					fmt.Println(s.Text())  //player number
					fmt.Println(playerSlice) //player pos and number
				})
			} else if i == 7 { // home team
				v.Find("font").Each(func(_ int, s *goquery.Selection) {
					playerStr, _ := s.Attr("title")
					playerSlice := utils.PlayerPosAndName(playerStr)
					fmt.Println(s.Text())  //player number
					fmt.Println(playerSlice) //player pos and number
				})
			} else if i == 3 { // [period time elapsed : period time remaining]
				fmt.Println(utils.PeriodClock(v.Text()))	
			} else { // event info
				fmt.Println(v.Text())
			}
		}
		fmt.Println("=====")
	})

	c.Visit("http://www.nhl.com/scores/htmlreports/20212022/PL020562.HTM") 
}