package main

import (
	"sportradar/score"
)

func main()  {

	teams := score.MatchStore{Store: []score.FootbalMatch{}}
	teams.AddMatch(0, 0,  "Mexico", "Canada")
	teams.AddMatch(0, 1,  "Spain", "Brazil")
	teams.AddMatch(0, 2, "Germany", "France")
	teams.AddMatch(0, 3, "Uruguay", "Italy")
	teams.AddMatch(0, 4, "Argentina", "Australia")

	teams.UpdateScore(1, "Spain", 2)
	teams.UpdateScore(1, "Spain", 2)
	teams.UpdateScore(1, "Spain", 2)
	teams.UpdateScore(0, "Mexico", 6)
	teams.UpdateScore(2, "Germany", 5)
	teams.UpdateScore(3, "Uruguay", 2)
	teams.UpdateScore(4, "Argentina", 10)

	teams.FinishMatch(3)
	//teams.FinishMatch(4)
	//teams.FinishMatch(1)

	teams.Sort()
}



