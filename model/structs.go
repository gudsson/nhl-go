package model

type Player struct {
	Position string
	Name string
	Number int8
}

type Event struct {
	EventId int16
	Period int8
	Strength string
	PeriodTimeElapsed string
	PeriodTimeRemaining string
	EventCode string
	EventDescription string
	AwayOnIce []Player
	HomeOnIce []Player
}