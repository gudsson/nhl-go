package model

import "time"

type Player struct {
	Position string
	Name string
	Number int
}

type Event struct {
	EventId int
	Period int
	Strength string
	PeriodTimeElapsed string
	PeriodTimeRemaining string
	EventCode string
	EventDescription string
	AwayOnIce []Player
	HomeOnIce []Player
	XCoordinate float32
	YCoordinate float32
}

type LiveFeedJSON struct {
	Copyright string `json:"copyright"`
	GamePk    int    `json:"gamePk"`
	Link      string `json:"link"`
	MetaData  struct {
		Wait      int    `json:"wait"`
		TimeStamp string `json:"timeStamp"`
	} `json:"metaData"`
	LiveData struct {
		Plays struct {
			AllPlays []struct {
				Players []struct {
					Player struct {
						ID       int    `json:"id"`
						FullName string `json:"fullName"`
						Link     string `json:"link"`
					} `json:"player"`
					PlayerType string `json:"playerType"`
				} `json:"players"`
				Result struct {
					Event       string `json:"event"`
					EventCode   string `json:"eventCode"`
					EventTypeID string `json:"eventTypeId"`
					Description string `json:"description"`
				} `json:"result"`
				About struct {
					EventIdx            int       `json:"eventIdx"`
					EventID             int       `json:"eventId"`
					Period              int       `json:"period"`
					PeriodType          string    `json:"periodType"`
					OrdinalNum          string    `json:"ordinalNum"`
					PeriodTime          string    `json:"periodTime"`
					PeriodTimeRemaining string    `json:"periodTimeRemaining"`
					DateTime            time.Time `json:"dateTime"`
					Goals               struct {
						Away int `json:"away"`
						Home int `json:"home"`
					} `json:"goals"`
				} `json:"about"`
				Coordinates struct {
					X float32 `json:"x"`
					Y float32 `json:"y"`
				} `json:"coordinates"`
				Team struct {
					ID      int    `json:"id"`
					Name    string `json:"name"`
					Link    string `json:"link"`
					TriCode string `json:"triCode"`
				} `json:"team"`
			} `json:"allPlays"`
		} `json:"plays"`
	} `json:"liveData"`
}

type AllPlays []struct {
	Players []struct {
		Player struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"player"`
		PlayerType string `json:"playerType"`
	} `json:"players"`
	Result struct {
		Event       string `json:"event"`
		EventCode   string `json:"eventCode"`
		EventTypeID string `json:"eventTypeId"`
		Description string `json:"description"`
	} `json:"result"`
	About struct {
		EventIdx            int       `json:"eventIdx"`
		EventID             int       `json:"eventId"`
		Period              int       `json:"period"`
		PeriodType          string    `json:"periodType"`
		OrdinalNum          string    `json:"ordinalNum"`
		PeriodTime          string    `json:"periodTime"`
		PeriodTimeRemaining string    `json:"periodTimeRemaining"`
		DateTime            time.Time `json:"dateTime"`
		Goals               struct {
			Away int `json:"away"`
			Home int `json:"home"`
		} `json:"goals"`
	} `json:"about"`
	Coordinates struct {
		X float32 `json:"x"`
		Y float32 `json:"y"`
	} `json:"coordinates"`
	Team struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Link    string `json:"link"`
		TriCode string `json:"triCode"`
	} `json:"team"`
}