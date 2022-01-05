package scrapers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gudsson/nhl-go/model"
)

func GetFeed(gameId string) model.AllPlays {
	resp, err := http.Get(fmt.Sprintf("https://statsapi.web.nhl.com/api/v1/game/%s/feed/live", gameId))
	if err != nil {
			fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var result model.LiveFeedJSON
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	return result.LiveData.Plays.AllPlays
}