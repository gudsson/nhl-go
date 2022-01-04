package scrapers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gudsson/nhl-go/model"
)

func GetFeed() model.AllPlays {
	resp, err := http.Get("https://statsapi.web.nhl.com/api/v1/game/2021020562/feed/live")
	if err != nil {
			fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var result model.LiveFeedJSON
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	// periodEvents := []model.AllPlays{}
	// inPeriod := false

	// for _, play := range result.LiveData.Plays.AllPlays {
	// 	// if play.Result.EventTypeID == "PERIOD_START" {
	// 	// 	inPeriod = true
	// 	// } else if play.Result.EventTypeID == "PERIOD_END" {
	// 	// 	inPeriod = false
	// 	// }
	// 	fmt.Printf("t1: %T\n", play)

	// 	// if inPeriod {
	// 	// 	periodEvents = append(periodEvents, play)
	// 	// }
	// }


	// fmt.Println(result.LiveData.Plays.AllPlays[9])
	return result.LiveData.Plays.AllPlays
}