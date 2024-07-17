package game

import "sync"

const STATUS_IN_PROGRESS = 0
const STATUS_FAIL = 5
const STATUS_WON = 10
const TYPE_5_WORDS = 5

type Game struct {
	ID     int64  `json:"id"`
	Puzzle string `json:"-"`
	UserId int64  `json:"user_id"`
	Status int8   `json:"status"`
	Type   int8   `json:"type"`
	Words  map[int8]*Word
}

func (g *Game) Calc() {
	wg := sync.WaitGroup{}
	wg.Add(len(g.Words))

	for _, word := range g.Words {
		go func() {
			word.CalcAgainst(g.Puzzle)
			wg.Done()
		}()
	}
	wg.Wait()
}

func (g *Game) InProgress() bool {
	return g.Status == STATUS_IN_PROGRESS
}

func (g *Game) AddWord(word string) {
	g.Words[int8(len(g.Words))] = &Word{Word: word}
}
