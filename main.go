package main

import (
	"fmt"
	"os"

	"github.com/gudsson/nhl-go/scrapers"
)

func main() {
	gameId := os.Args[1]
	events := scrapers.CombinePbpAndFeed(gameId)

	for _, event := range events {
		fmt.Printf("%d. %s, [%f, %f]\n", event.EventId, event.EventCode, event.XCoordinate, event.YCoordinate)
	}
}