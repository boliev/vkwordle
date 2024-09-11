package pg

import (
	"database/sql"
	"github.com/boliev/vkwordle/internal/domain/game"
)

type Puzzle struct {
	DB *sql.DB
}

func NewPuzzle(db *sql.DB) *Puzzle {
	return &Puzzle{
		DB: db,
	}
}

func (p *Puzzle) GetRandomPuzzle() (*game.Puzzle, error) {
	const s = "SELECT word, category, hint FROM puzzles ORDER BY RANDOM() LIMIT 1"

	puzzle := &game.Puzzle{}
	row := p.DB.QueryRow(s)

	err := row.Scan(&puzzle.Word, &puzzle.Category, &puzzle.Hint)
	if err != nil {
		return nil, err
	}

	return puzzle, nil
}

func (p *Puzzle) GetPuzzle(word string) (*game.Puzzle, error) {
	const s = "SELECT word, category, hint FROM puzzles WHERE word = $1"

	puzzle := &game.Puzzle{}
	row := p.DB.QueryRow(s, word)

	err := row.Scan(&puzzle.Word, &puzzle.Category, &puzzle.Hint)
	if err != nil {
		return nil, err
	}

	return puzzle, nil
}
