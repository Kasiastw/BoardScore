package main

import (
	"log"
	"sync/atomic"
)

type (
	footbalMatch struct {
		id int64
		idGame int
		game map[string]int
		sum int
	}
	MatchStore struct {
		store []footbalMatch
	}
)

var id int64

func (fs MatchStore) NewMatch(id int64, idGame int, hometeam string, awayteam string) footbalMatch {
	log.Printf("The match: %s vs %s started", hometeam, awayteam)
	return footbalMatch{
		id: id,
		idGame: idGame,
		game: map[string]int{hometeam:0, awayteam:0},
		sum: 0,
	}
}

func (fs MatchStore) updateScore(idGame int, team string, score int) []footbalMatch {
	count:= 0
	for i :=range fs.store{
		if fs.store[i].idGame == idGame {
			if _, ok := fs.store[i].game[team]; ok {
				id = atomic.AddInt64(&id, 1)
				fs.store[i].id = id
				fs.store[i].game[team] = fs.store[i].game[team] + score
				fs.store[i].sum = fs.store[i].sum + score
				count=1
				log.Printf("A gol for %s => The current score is %v", team, fs.store[i].game)
			} else {
				log.Printf("[ERROR] There is no team %s with idGame= %v", team, idGame)
			}
		}
	}
	if count==0{
		log.Printf("[ERROR] There is no match with idGame=%v", idGame)
	}
	return fs.store
}

func (fs *MatchStore) finishMatch(idGame int) []footbalMatch {
	count:= 0
	for i := 0; i < len(fs.store); i++ {
		if fs.store[i].idGame==idGame {
			fs.store = append(fs.store[:i],fs.store[i+1:]...)
			count++
			log.Printf("The match with id = %v has been finished", idGame)
			break
		}
	}
	if count==0{
		log.Printf("[ERROR] There is no match to finish with idGame=%v", idGame)
	}
	return fs.store
}

func (fs MatchStore) sort() []footbalMatch{
	ok := false
	for !ok {
		ok = true
		i:=0
		for i<len(fs.store)-1 {

			if fs.store[i].sum < fs.store[i + 1].sum && fs.store[i].sum != fs.store[i+1].sum {
				fs.store[i], fs.store[i + 1] = fs.store[i + 1], fs.store[i]
				ok = false
				}
			if fs.store[i].sum == fs.store[i+1].sum && fs.store[i].id < fs.store[i + 1].id{
				fs.store[i], fs.store[i + 1] = fs.store[i + 1], fs.store[i]
				ok = false
				}
			i++
		}
	}
	log.Println("sorted board score", fs.store)
	return fs.store
	}

func main()  {

	teams := MatchStore{store: []footbalMatch{}}
	teams.store = append(teams.store, teams.NewMatch(0, 0,  "Mexico", "Canada"))
	teams.store = append(teams.store, teams.NewMatch(0, 1,  "Spain", "Brazil"))
	teams.store = append(teams.store, teams.NewMatch(0, 2, "Germany", "France"))
	teams.store = append(teams.store, teams.NewMatch(0, 3, "Uruguay", "Italy"))
	teams.store = append(teams.store, teams.NewMatch(0, 4, "Argentina", "Australia"))

	teams.updateScore(1, "Spain", 2)
	teams.updateScore(1, "Spain", 2)
	teams.updateScore(1, "Spain", 2)
	teams.updateScore(0, "Mexico", 6)
	teams.updateScore(2, "Germany", 5)
	teams.updateScore(3, "Uruguay", 2)
	teams.updateScore(4, "Argentina", 10)

	teams.finishMatch(3)
	//teams.finishMatch(4)
	//teams.finishMatch(1)

	teams.sort()
}



