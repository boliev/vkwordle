package game

import (
	"fmt"
	"sync"
)

const STATUS_IN_PROGRESS = 0
const STATUS_FAIL = 5
const STATUS_WON = 10
const TYPE_5_WORDS = 5

type Game struct {
	ID     int64  `json:"id"`
	Puzzle string `json:"-"`
	UserId int64  `json:"user_id"`
	Status int    `json:"status"`
	Type   int    `json:"type"`
	Words  map[int]*Word
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
	g.Words[len(g.Words)] = &Word{Word: word}
}

func (g *Game) WordsLen() int {
	return len(g.Words)
}

func (g *Game) lastWord() (*Word, error) {
	if len(g.Words) == 0 {
		return nil, fmt.Errorf("game has no words")
	}

	return g.Words[g.WordsLen()-1], nil
}

func (g *Game) IsLastWordCorrect() bool {
	if len(g.Words) == 0 {
		return false
	}
	lastWord, err := g.lastWord()
	if err != nil {
		return false
	}

	for _, status := range lastWord.Letters {
		if status != LETTER_STATE_IN_PLACE {
			return false
		}
	}
	return true
}
