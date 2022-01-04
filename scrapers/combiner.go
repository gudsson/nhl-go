package scrapers

import (
	"fmt"

	"github.com/gudsson/nhl-go/model"
	"github.com/gudsson/nhl-go/utils"
)

func CombinePbpAndFeed() []model.Event {
	feedEvents := GetFeed()
	pbpEvents := GetEvents()
	inFeedPeriod := false
	inPbpPeriod := false

	feedIdx := 0
	pbpIdx := 0


	for pbpIdx < len(pbpEvents) {
		feedCode := feedEvents[feedIdx].Result.EventTypeID
		pbpCode := pbpEvents[pbpIdx].EventCode

		if feedCode == "PERIOD_START" {
			inFeedPeriod = true
		} else if feedCode == "PERIOD_END" {
			inFeedPeriod = false
		}

		if pbpCode == "PSTR" {
			inPbpPeriod = true
		} else if pbpCode == "PEND" {
			inPbpPeriod = false
		}

		if (inFeedPeriod || (feedCode == "PERIOD_END")) && (inPbpPeriod || (pbpCode == "PEND")) {
			feedEvent, pbpEvent := feedEvents[feedIdx], &pbpEvents[pbpIdx]

			pbpTime := utils.ConvertPeriodTime(pbpEvent.PeriodTimeElapsed)
			feedTime := utils.ConvertPeriodTime(feedEvent.About.PeriodTime)

			if (feedEvent.About.Period == pbpEvent.Period) && (pbpTime == feedTime) {
				pbpEvent.XCoordinate = feedEvent.Coordinates.X
				pbpEvent.YCoordinate = feedEvent.Coordinates.Y

				feedIdx++
				pbpIdx++
			} else {
				if (feedTime > pbpTime) {
					pbpIdx++
				} else {
					feedIdx++
					fmt.Println("extra feed event and I don't know why")
				}
			}
		} else if !inPbpPeriod && !inFeedPeriod{
			feedIdx++
			pbpIdx++
		} else if !inFeedPeriod {
			feedIdx++
		} else if !inPbpPeriod {
			pbpIdx++
		}
	}
	return pbpEvents
}