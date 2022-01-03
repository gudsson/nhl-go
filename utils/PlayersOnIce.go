package utils

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gudsson/nhl-go/model"
)

func PlayersOnIce(tds *goquery.Selection) []model.Player {
	players := []model.Player{}
	tds.Find("font").Each(func(_ int, s *goquery.Selection) {
		playerStr, _ := s.Attr("title")
		playerSlice := PlayerPosAndName(playerStr)
		playerNum, _ := strconv.Atoi(s.Text())

		players = append(
			players,
			model.Player{Position: playerSlice[0], Name: playerSlice[1], Number: int8(playerNum)},
		)
	})

	return players
}