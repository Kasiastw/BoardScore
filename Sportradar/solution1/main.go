package main

import (
	"./score"
	"log"
	"sort"
)

func main()  {

	teams := score.ScoreBoard{Board: map[int]score.FootbalMatch{}}
	teams.AddMatch("Mexico", "Canada")
	teams.AddMatch("Spain", "Brazil")
	teams.AddMatch("Germany", "France")
	teams.AddMatch("Uruguay", "Italy")
	teams.AddMatch("Argentina", "Australia")
	log.Println(teams)

	teams.UpdateScore(1, "Spain", "Brazil",  2, 0)
	teams.UpdateScore(1, "Spain", "Brazil",  2, 0)
	teams.UpdateScore(1, "Spain", "Brazil",  2, 0)
	teams.UpdateScore(0, "Mexico", "Canada", 6, 0)
	teams.UpdateScore(2, "Germany", "France", 5, 0)
	teams.UpdateScore(3, "Uruguay", "Italy", 2, 0)
	teams.UpdateScore(4, "Argentina", "Australia", 10, 0)

	log.Println(teams)
	//teams.FinishMatch(3)
	//teams.FinishMatch(4)
	//teams.FinishMatch(1)

	sort.Sort(teams)
	log.Println("sorted board score: ", teams)
}
