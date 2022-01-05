# nhl-go

Golang scraper (using Colly) for NHL data.

The NHL provides multiple sets of data for events that occurs throughout a game. 

Play-by-Play data has been provided since 2006, and contains records of every player on the ice for a specific event, which has spawned "advanced" metrics such as Corsi and Fenwick, which track shot attempts a given player was on-ice for.

Since partnering with MLB Advanced media, the NHL has provided X-Y on-ice coordinates for these events...under different event IDs.  These coordinates (again, under different IDs for some reason) don't contain records of who was on ice for the event.

This scraper combines Play-by-Play data with the MLB advanced media data (which provides X-Y coordinates), enabling the creation of Expected Goal (xG) models, which is the next step for this repo.

To get event data augmented with X-Y Coordinates, run main.go and provide the desired NHL gameId

```
git clone https://github.com/gudsson/nhl-go.git
cd nhl-go

# example gameId 2021020562: EDM 5 - NJD 6 (OT) on December 31, 2021
go run main.go 2021020562
```

