package score

import (
	"reflect"
	"testing"
)

func (fs MatchStore) updateScore(idGame int, team string, score int) []FootbalMatch {
	for i :=range fs.Store{
		if fs.Store[i].idGame == idGame {
			if _, ok := fs.Store[i].game[team]; ok {
				fs.Store[i].game[team] = fs.Store[i].game[team] + score
				fs.Store[i].sum = fs.Store[i].sum + score
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
		team   string
		score  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []FootbalMatch
	}{
		{
			name: "test if the score was updated when team and idGame exist",
			fields: fields{store: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}}},
			args: args{idGame: 1, score: 2, team: "A"},
			want: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 2, "B":0}, sum:2},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
		{
			name: "test if the score wasn't updated because idGame doesn't exist",
			fields: fields{store: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}}},
			args: args{idGame: 111, score: 2, team: "A"},
			want: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
		{
			name: "test if the score wasn't updated because team doesn't exist",
			fields: fields{store: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}}},
			args: args{idGame: 1, score: 2, team: "C"},
			want: []FootbalMatch{FootbalMatch{id: 0, idGame: 1, game: map[string]int{"A": 0, "B":0}, sum:0},
				FootbalMatch{id: 0, idGame: 22, game: map[string]int{"A": 0, "B":0}, sum:0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := MatchStore{
				Store: tt.fields.store,
			}
			if got := fs.updateScore(tt.args.idGame, tt.args.team, tt.args.score); !reflect.DeepEqual(got, tt.want) {
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


