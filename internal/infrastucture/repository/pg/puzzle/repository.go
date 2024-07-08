package pg_repository

import (
	"database/sql"
	domain "github.com/boliev/vkwordle/internal/domain/puzzle"
)

type Puzzle struct {
	DB *sql.DB
}

func NewPuzzle(db *sql.DB) *Puzzle {
	return &Puzzle{
		DB: db,
	}
}

func (p *Puzzle) GetPuzzle() *domain.Puzzle {
	return &domain.Puzzle{
		Word:     "пират",
		Category: "криминал",
		Hint:     "Морской волк",
	}
}
