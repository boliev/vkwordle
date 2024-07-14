package game

type PuzzleRepository interface {
	GetRandomPuzzle() (*Puzzle, error)
}

type GameRepository interface {
	CreateGame(userId int64, word string) (*Game, error)
	GetActiveGame(userId int64, gameType int8) (*Game, error)
}
