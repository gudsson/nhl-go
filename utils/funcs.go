package utils

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gudsson/nhl-go/model"
)

func ConvertPeriodTime(periodTime string) int {
	time, _ := strconv.Atoi(strings.Replace(periodTime, ":", "", -1))
	return time
}

func PlayersOnIce(tds *goquery.Selection) []model.Player {
	players := []model.Player{}
	tds.Find("font").Each(func(_ int, s *goquery.Selection) {
		playerStr, _ := s.Attr("title")
		playerSlice := PlayerPosAndName(playerStr)
		playerNum, _ := strconv.Atoi(s.Text())

		players = append(
			players,
			model.Player{Position: playerSlice[0], Name: playerSlice[1], Number: int(playerNum)},
		)
	})

	return players
}

func PlayerPosAndName(str string) []string {
	return strings.Split(str, " - ")
}

func PeriodClock(str string) []string {
	re := regexp.MustCompile(`(\d{1,2}:\d{2})`)
	split := re.FindAllString(str, 2)
	times := []string{}
	for i := range split {
		times = append(times, split[i])
	}

	return times
}