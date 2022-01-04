package main

// PBP: http://www.nhl.com/scores/htmlreports/20212022/PL020562.HTM
// live: https://statsapi.web.nhl.com/api/v1/game/2021020562/feed/live

import (
	"fmt"

	"github.com/gudsson/nhl-go/scrapers"
)

func main() {
	events := scrapers.CombinePbpAndFeed()

	for _, event := range events {
		fmt.Printf("%d. %s, [%f, %f]\n", event.EventId, event.EventCode, event.XCoordinate, event.YCoordinate)
	}
}