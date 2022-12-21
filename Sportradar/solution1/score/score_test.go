package score

import (
	"reflect"
	"sort"
	"testing"
)

func (sc ScoreBoard) updateScore(idGame int, hometeam string, awayteam string, scoreHomeTeam int, scoreAwayTeam int ) map[int]FootbalMatch {
	if _, ok:= sc.Board[idGame]; ok {
		if _, ok := sc.Board[idGame].match[hometeam]; ok {
			if _, ok := sc.Board[idGame].match[awayteam]; ok {
				sc.Board[idGame].match[hometeam] += scoreHomeTeam
				sc.Board[idGame].match[awayteam] += scoreAwayTeam
				sc.Board[idGame].match["sum"] += scoreHomeTeam + scoreAwayTeam
			}
		}
	}
	return sc.Board
}

func TestScoreBoard_updateScore(t *testing.T) {
	type (
		args struct {
			idGame int
			hometeam string
			awayteam string
			scoreHomeTeam int
			scoreAwayTeam int
		}
		fields struct {
			board map[int]FootbalMatch
		}
	)

	tests := []struct{
		name 	string
		fields 	fields
		args	args
		want	map[int]FootbalMatch
	}{
		{
			name: "test if the score was updated when both teams and idGame exist",
			fields: fields{board: map[int]FootbalMatch{
				0:{id: 0, match: map[string]int{"Mexico":0, "Canada":0, "sum":0}},
				1:{id: 1, match: map[string]int{"Brazil":0, "Germany":0, "sum":0}}}},
			args: args{idGame: 1, hometeam:"Brazil", awayteam:"Germany", scoreHomeTeam: 2, scoreAwayTeam: 0},
			want: map[int]FootbalMatch{
				0:{id: 0, match: map[string]int{"Mexico":0, "Canada":0, "sum":0}},
				1:{id: 1, match: map[string]int{"Brazil":2, "Germany":0, "sum":2}}},
		},
		{
			name: "test if the score wasn't updated because idGame doesn't exist",
			fields: fields{board: map[int]FootbalMatch{
				0:{id: 0, match: map[string]int{"Mexico":0, "Canada":0, "sum":0}},
				1:{id: 1, match: map[string]int{"Brazil":0, "Germany":0, "sum":0}}}},
			args: args{idGame: 111, hometeam:"Brazil", awayteam:"Germany", scoreHomeTeam: 2, scoreAwayTeam: 0},
			want: map[int]FootbalMatch{
				0:{id: 0, match: map[string]int{"Mexico":0, "Canada":0, "sum":0}},
				1:{id: 1, match: map[string]int{"Brazil":0, "Germany":0, "sum":0}}},
		},
		{
			name: "test if the score wasn't updated because the team doesn't exist",
			fields: fields{board: map[int]FootbalMatch{
				0:{id: 0, match: map[string]int{"Mexico":0, "Canada":0, "sum":0}},
				1:{id: 1, match: map[string]int{"Brazil":0, "Germany":0, "sum":0}}}},
			args: args{idGame: 1, hometeam:"Spain", awayteam:"Germany", scoreHomeTeam: 2, scoreAwayTeam: 0},
			want: map[int]FootbalMatch{
				0:{id: 0, match: map[string]int{"Mexico":0, "Canada":0, "sum":0}},
				1:{id: 1, match: map[string]int{"Brazil":0, "Germany":0, "sum":0}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := ScoreBoard{
				Board: tt.fields.board,
			}
			if got := sc.updateScore(tt.args.idGame, tt.args.hometeam, tt.args.awayteam, tt.args.scoreHomeTeam, tt.args.scoreAwayTeam); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {

	type (
		fields struct {
			board map[int]FootbalMatch
		}
	)

	tests := []struct{
		name 	string
		fields 	fields
		want	fields
	}{
		{
			name: "test if the map was sorted",
			fields: fields{board: map[int]FootbalMatch{
				0:{0, map[string]int{"Canada":0, "Mexico":6, "sum":6}},
				1:{1, map[string]int{"Brazil":0, "Spain":6, "sum":6}},
				2:{2, map[string]int{"France":0, "Germany":5, "sum":5}},
				3:{3, map[string]int{"Italy":0, "Uruguay":2, "sum":2}},
				4:{4, map[string]int{"Argentina":10, "Australia":0, "sum":10}}}},
			want: fields{board: map[int]FootbalMatch{
				0:{4, map[string]int{"Argentina":10, "Australia":0, "sum":10}},
				1:{0, map[string]int{"Canada":0, "Mexico":6, "sum":6}},
				2:{1, map[string]int{"Brazil":0, "Spain":6, "sum":6}},
				3:{2, map[string]int{"France":0, "Germany":5, "sum":5}},
				4:{3, map[string]int{"Italy":0, "Uruguay":2, "sum":2}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := ScoreBoard{
				Board: tt.fields.board,
			}
			if sort.Sort(&sc); reflect.DeepEqual(sc, tt.want) {
				t.Errorf("sort score board() = %v, \n want %v", sc, tt.want)
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	sc := ScoreBoard{Board: map[int]FootbalMatch{
		0:{0, map[string]int{"Canada":0, "Mexico":6, "sum":6}},
		1:{1, map[string]int{"Brazil":0, "Spain":6, "sum":6}},
		2:{2, map[string]int{"France":0, "Germany":5, "sum":5}},
		3:{3, map[string]int{"Italy":0, "Uruguay":2, "sum":2}},
		4:{4, map[string]int{"Argentina":10, "Australia":0, "sum":10}}}}
	b.ResetTimer()
	sort.Sort(sc)
}





