package score

import "log"

type (
	FootbalMatch struct {
		id int64
		idGame int
		game map[string]int
		sum int
	}
	MatchStore struct {
		Store []FootbalMatch
	}
)

var id int64

func NewMatch(id int64, idGame int, hometeam string, awayteam string) FootbalMatch {
	log.Printf("The match: %s vs %s started", hometeam, awayteam)
	return FootbalMatch{
		id: id,
		idGame: idGame,
		game: map[string]int{hometeam:0, awayteam:0},
		sum: 0,
	}
}
func (fs *MatchStore) AddMatch(id int64, idGame int, hometeam string, awayteam string) {
	fs.Store = append(fs.Store, NewMatch(id, idGame, hometeam, awayteam))
}

func (fs MatchStore) UpdateScore(idGame int, hometeam string, awayteam string, scoreHomeTeam int, scoreAwayTeam int ) {
	count:= 0
	for i :=range fs.Store {
		if fs.Store[i].idGame == idGame {
			if _, ok := fs.Store[i].game[hometeam]; ok {
				if _, ok := fs.Store[i].game[awayteam]; ok {
					id++
					fs.Store[i].id = id
					fs.Store[i].game[hometeam] = fs.Store[i].game[hometeam] + scoreHomeTeam
					fs.Store[i].game[awayteam] = fs.Store[i].game[awayteam] + scoreAwayTeam
					fs.Store[i].sum = fs.Store[i].sum + scoreHomeTeam + scoreAwayTeam
					count++
					log.Printf("A gol for %s and %s  => The current score is %v", hometeam, awayteam, fs.Store[i].game)
				} else {
					log.Printf("[ERROR] There is no hometeam %s with idGame= %v", awayteam, idGame)
				}
			} else {
				log.Printf("[ERROR] There is no team %s with idGame= %v", hometeam, idGame)
			}
		}
	}
	if count==0{
		log.Printf("[ERROR] There is no match with idGame=%v", idGame)
	}
}

func (fs *MatchStore) FinishMatch(idGame int) {
	count:= 0
	for i := 0; i < len(fs.Store); i++ {
		if fs.Store[i].idGame==idGame {
			fs.Store = append(fs.Store[:i],fs.Store[i+1:]...)
			count++
			log.Printf("The match with id = %v has been finished", idGame)
			break
		}
	}
	if count==0{
		log.Printf("[ERROR] There is no match to finish with idGame=%v", idGame)
	}
}

func (fs MatchStore) Sort() {
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
	log.Println("sorted board score", fs.Store)
}

