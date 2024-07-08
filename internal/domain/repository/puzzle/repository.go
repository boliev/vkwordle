package domain

import domain "github.com/boliev/vkwordle/internal/domain/puzzle"

type PuzzleRepository interface {
	GetPuzzle() *domain.Puzzle
}
