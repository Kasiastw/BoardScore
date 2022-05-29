package main

import (
	"reflect"
	"testing"
)

type (
	scoreUpdate struct {
		idGame int
		score int
		team string
	}
)

var mStore = MatchStore{store: []footbalMatch{}}

var testCasesUpdateScore = []struct{
	description string
	input scoreUpdate
	expected []footbalMatch
}{
	{
		description: "test if the score wasn't updated because idGame doesn't exist",
		input: scoreUpdate{idGame: 10, score: 2, team: "A"},
		expected: []footbalMatch{footbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
					footbalMatch{id: 0, idGame: 11, game: map[string]int{"A": 0, "B":0}, sum:0},
					footbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
	{
		description: "test if the score wasn't updated because team doesn't exist",
		input: scoreUpdate{idGame: 1, score: 2, team: "C"},
		expected: []footbalMatch{footbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
					footbalMatch{id: 0, idGame: 11, game: map[string]int{"A": 0, "B":0}, sum:0},
					footbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
	{
		description: "test if the score was updated when team and idGame exist",
		input: scoreUpdate{idGame: 1, score: 2, team: "A"},
		expected: []footbalMatch{footbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 2, "B":0}, sum:2},
					footbalMatch{id: 0, idGame: 11, game: map[string]int{"A": 0, "B":0}, sum:0},
					footbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
	}

var testCasesFinishMatch = []struct{
	description string
	idGame int
	expected []footbalMatch
}{
	{
		description: "test if the match was deleted when idGame exists",
		idGame: 11,
		expected: []footbalMatch{footbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 2, "B":0}, sum:2},
					footbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
	{
		description: "test if the match wasn't deleted because idGame doesn't exist",
		idGame: 100,
		expected: []footbalMatch{
				footbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 2, "B":0}, sum:2},
				footbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
	}

var testCasesSort = []struct{
	description string
	input MatchStore
	expected []footbalMatch
}{
	{
		description: "there are no games with the same total score",
		input: MatchStore{store: []footbalMatch{footbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 2, "B":0}, sum:2},
			footbalMatch{id: 2, idGame: 11, game: map[string]int{"A": 3, "B":0}, sum:3},
			footbalMatch{id: 3, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}}},
		expected: []footbalMatch{footbalMatch{id: 2, idGame: 11, game: map[string]int{"A": 3, "B":0}, sum:3},
			footbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 2, "B":0}, sum:2},
			footbalMatch{id: 3, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
	},
	{
		description: "there are the games with the same total score",
		input: MatchStore{store: []footbalMatch{footbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 2, "B":0}, sum:2},
			footbalMatch{id: 2, idGame: 11, game: map[string]int{"A": 2, "B":0}, sum:2},
			footbalMatch{id: 3, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}}},
		expected: []footbalMatch{footbalMatch{id: 2, idGame: 11, game: map[string]int{"A": 2, "B":0}, sum:2},
			footbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 2, "B":0}, sum:2},
			footbalMatch{id: 3, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
	},
}

func (fs MatchStore) NewMatch(id int64, idGame int, hometeam string, awayteam string) footbalMatch {
	return footbalMatch{
		id: id,
		idGame: idGame,
		game: map[string]int{hometeam:0, awayteam:0},
		sum: 0,
	}
}

func (fs MatchStore) updateScoreT(idGame int, team string, score int) []footbalMatch {
	for i :=range fs.store{
		if fs.store[i].idGame == idGame {
			if _, ok := fs.store[i].game[team]; ok {
				fs.store[i].id = 0
				fs.store[i].game[team] = fs.store[i].game[team] + score
				fs.store[i].sum = fs.store[i].sum + score
			}
		}
	}
	return fs.store
}

func (fs *MatchStore) finishMatchT(idGame int) []footbalMatch {
	count:= 0
	for i := 0; i < len(fs.store); i++ {
		if fs.store[i].idGame==idGame {
			fs.store = append(fs.store[:i],fs.store[i+1:]...)
			count++
			break
		}
	}
	if count==0{
		return fs.store
	}
	return fs.store
}

func TestNewMatch(t *testing.T)  {
	mStore.store = append(mStore.store, mStore.NewMatch(0, 1, "A", "B"))
	mStore.store = append(mStore.store, mStore.NewMatch(0, 11, "A", "B"))
	mStore.store = append(mStore.store, mStore.NewMatch(0, 22, "A", "B"))
}
func TestUpdateScore(t *testing.T) {
	for _, tt := range testCasesUpdateScore {
		got:=  mStore.updateScoreT(tt.input.idGame, tt.input.team, tt.input.score)
		want:= tt.expected

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestFinishMatch(t *testing.T) {
	for _, tt := range testCasesFinishMatch {
		got:=  mStore.finishMatchT(tt.idGame)
		want:= tt.expected

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestSort(t *testing.T)  {
	for _, tt := range testCasesSort {
		mStore = tt.input
		got := mStore.sortT()
		want:= tt.expected

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func (fs MatchStore) sortT() []footbalMatch{
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
	return fs.store
}

