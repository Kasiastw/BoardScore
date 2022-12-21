package score

import (
	"reflect"
	"testing"
)

func (fs MatchStore) updateScore(idGame int, hometeam string, awayteam string, scoreHomeTeam int, scoreAwayTeam int) []FootbalMatch {
	for i :=range fs.Store {
		if fs.Store[i].idGame == idGame {
			if _, ok := fs.Store[i].game[hometeam]; ok {
				if _, ok := fs.Store[i].game[awayteam]; ok {
					fs.Store[i].game[hometeam] = fs.Store[i].game[hometeam] + scoreHomeTeam
					fs.Store[i].game[awayteam] = fs.Store[i].game[awayteam] + scoreAwayTeam
					fs.Store[i].sum = fs.Store[i].sum + scoreHomeTeam + scoreAwayTeam
				}
			}
		}
	}
	return fs.Store
}

func (fs *MatchStore) finishMatch(idGame int) []FootbalMatch {
	for i := 0; i < len(fs.Store); i++ {
		if fs.Store[i].idGame==idGame {
			fs.Store = append(fs.Store[:i],fs.Store[i+1:]...)
			break
		}
	}
	return fs.Store
}

func (fs MatchStore) sort() []FootbalMatch{
	ok := false
	for !ok {
		ok = true
		i:=0
		for i<len(fs.Store)-1 {
			if fs.Store[i].sum < fs.Store[i + 1].sum && fs.Store[i].sum != fs.Store[i+1].sum {
				fs.Store[i], fs.Store[i + 1] = fs.Store[i + 1], fs.Store[i]
				ok = false
			}
			if fs.Store[i].sum == fs.Store[i+1].sum && fs.Store[i].id < fs.Store[i + 1].id{
				fs.Store[i], fs.Store[i + 1] = fs.Store[i + 1], fs.Store[i]
				ok = false
			}
			i++
		}
	}
	return fs.Store
}

func TestMatchStore_updateScore(t *testing.T) {
	type fields struct {
		store []FootbalMatch
	}
	type args struct {
		idGame int
		hometeam string
		awayteam string
		scoreHomeTeam int
		scoreAwayTeam int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []FootbalMatch
	}{
		{
			name: "test if the score was updated when both teams and idGame exist",
			fields: fields{store: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}}},
			args: args{idGame: 1, hometeam:"A", awayteam:"B", scoreHomeTeam: 2, scoreAwayTeam: 0},
			want: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 2, "B":0}, sum:2},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
		{
			name: "test if the score wasn't updated because idGame doesn't exist",
			fields: fields{store: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}}},
			args: args{idGame: 111, hometeam:"A", awayteam:"B", scoreHomeTeam: 2, scoreAwayTeam: 0},
			want: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
		{
			name: "test if the score wasn't updated because team doesn't exist",
			fields: fields{store: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}}},
			args: args{idGame: 1, hometeam:"C", awayteam:"B", scoreHomeTeam: 2, scoreAwayTeam: 0},
			want: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := MatchStore{
				Store: tt.fields.store,
			}
			if got := fs.updateScore(tt.args.idGame, tt.args.hometeam, tt.args.awayteam, tt.args.scoreHomeTeam, tt.args.scoreAwayTeam); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchStore_finishMatch(t *testing.T) {
	type fields struct {
		store []FootbalMatch
	}
	type args struct {
		idGame int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []FootbalMatch
	}{
		{
			name: "test if the match was deleted when idGame exists",
			fields: fields{store: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 11, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}}},
			args: args{idGame: 11},
			want: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
		{
			name: "test if the match wasn't deleted because doesn't idGame",
			fields: fields{store: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}}},
			args: args{idGame: 11},
			want: 	[]FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := MatchStore{Store: tt.fields.store}
			if got := fs.finishMatch(tt.args.idGame); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("finishMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchStore_sort(t *testing.T) {
	type fields struct {
		store []FootbalMatch
	}
	tests := []struct {
		name   string
		fields fields
		want   []FootbalMatch
	}{
		{
			name: "test if the games with the same total score are returned ordered by id",
			fields: fields{store: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 2, "B":0}, sum:2},
				FootbalMatch{id: 2, idGame: 11, game: map[string]int{"A": 2, "B":0}, sum:2},
				FootbalMatch{id: 3, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}}},
			want: []FootbalMatch{FootbalMatch{id: 2, idGame: 11, game: map[string]int{"A": 2, "B":0}, sum:2},
				FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 2, "B":0}, sum:2},
				FootbalMatch{id: 3, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
		{
			name: "test if the games with different total score are sorted",
			fields: fields{store: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 3, "B":0}, sum:3},
				FootbalMatch{id: 2, idGame: 11, game: map[string]int{"A": 2, "B":0}, sum:2},
				FootbalMatch{id: 3, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}}},
			want: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 3, "B":0}, sum:3},
				FootbalMatch{id: 2, idGame: 11, game: map[string]int{"A": 2, "B":0}, sum:2},
				FootbalMatch{id: 3, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := MatchStore{
				Store: tt.fields.store,
			}
			if got := fs.sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//func BenchmarkUpdateScore(b *testing.B) {
//	fm := MatchStore{Store: []FootbalMatch{
//		{4, 0, map[string]int{"Canada":0, "Mexico":6}, 6},
//		{3, 1, map[string]int{"Brazil":0, "Spain":6}, 6},
//		{5, 2, map[string]int{"France":0, "Germany":5}, 5},
//		{6, 2, map[string]int{"Italy":0, "Uruguay":2}, 2},
//		{7, 4, map[string]int{"Argentina":0, "Australia":2}, 2},
//	}}
//	b.ResetTimer()
//	fm.updateScore(0, "Argentina", "Australia", 0, 4)
//}

func BenchmarkUpdateScorePointer(b *testing.B) {
	fm := MatchStore{Store: []FootbalMatch{
		{4, 0, map[string]int{"Canada":0, "Mexico":6}, 6},
		{3, 1, map[string]int{"Brazil":0, "Spain":6}, 6},
		{5, 2, map[string]int{"France":0, "Germany":5}, 5},
		{6, 2, map[string]int{"Italy":0, "Uruguay":2}, 2},
		{7, 4, map[string]int{"Argentina":0, "Australia":2}, 2},
	}}
	b.ResetTimer()
	fm.updateScorePointer(0, "Argentina", "Australia", 0, 4)
}

func (fs *MatchStore) updateScorePointer(idGame int, hometeam string, awayteam string, scoreHomeTeam int, scoreAwayTeam int) []FootbalMatch {
	for i :=range fs.Store {
		if fs.Store[i].idGame == idGame {
			if _, ok := fs.Store[i].game[hometeam]; ok {
				if _, ok := fs.Store[i].game[awayteam]; ok {
					fs.Store[i].game[hometeam] = fs.Store[i].game[hometeam] + scoreHomeTeam
					fs.Store[i].game[awayteam] = fs.Store[i].game[awayteam] + scoreAwayTeam
					fs.Store[i].sum = fs.Store[i].sum + scoreHomeTeam + scoreAwayTeam
				}
			}
		}
	}
	return fs.Store
}

