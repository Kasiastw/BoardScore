package score

import (
	"log"
)

type (
	FootbalMatch struct {
		id int
		match map[string]int
	}
	ScoreBoard struct {
		Board map[int]FootbalMatch
	}
)

var idx int

func (sc ScoreBoard) AddMatch(hometeam string, awayteam string) {
	idx++
	idMatch := idx-1
	sc.Board[idMatch] = FootbalMatch{id: idMatch, match: map[string]int{hometeam:0, awayteam:0, "sum":0}}
	log.Printf("The match %s vs %s has been started", hometeam, awayteam)
}

func (sc ScoreBoard) UpdateScore(idGame int, hometeam string, awayteam string, scoreHomeTeam int, scoreAwayTeam int ) {
	if _, ok:= sc.Board[idGame]; ok {
		if _, ok := sc.Board[idGame].match[hometeam]; ok {
			if _, ok := sc.Board[idGame].match[awayteam]; ok {
				sc.Board[idGame].match[hometeam] += scoreHomeTeam
				sc.Board[idGame].match[awayteam] += scoreAwayTeam
				sc.Board[idGame].match["sum"] += scoreHomeTeam + scoreAwayTeam
				log.Printf("The current score for the match: %v", sc.Board[idGame].match)
			}else {
				log.Printf("[ERROR] There is no match (id=%v) with awayteam %s", idGame, awayteam)}
		} else {
			log.Printf("[ERROR] There is match (id=%v)with hometeam %s", idGame, hometeam)
		}
	} else {
		log.Printf("[ERROR] There is no match with id= %v", idGame)
	}
}

func (sc ScoreBoard) FinishMatch(idGame int) {
	if _, ok := sc.Board[idGame]; ok {
		delete(sc.Board, idGame)
		log.Printf("The match with id = %v has been finished", idGame)
	} else {
		log.Printf("[ERROR] There is no match to finish with idGame=%v", idGame)
	}
}

func (sc ScoreBoard) Len() int { return len(sc.Board)}

func (sc ScoreBoard) Less(i, j int) bool {
	return sc.Board[i].match["sum"] > sc.Board[j].match["sum"] || (sc.Board[i].match["sum"] == sc.Board[j].match["sum"] && sc.Board[i].id<sc.Board[j].id)
}

func (sc ScoreBoard) Swap(i, j int) { sc.Board[i], sc.Board[j] = sc.Board[j], sc.Board[i] }
