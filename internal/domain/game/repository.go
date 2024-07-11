package game

type PuzzleRepository interface {
	GetPuzzle() (*Puzzle, error)
}
