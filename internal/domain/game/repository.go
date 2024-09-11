package game

type PuzzleRepository interface {
	GetRandomPuzzle() (*Puzzle, error)
	GetPuzzle(word string) (*Puzzle, error)
}

type GameRepository interface {
	CreateGame(userId int64, word string) (*Game, error)
	GetActiveGame(userId int64, gameType int) (*Game, error)
	AddWord(gameId int64, word string) error
	Finish(gameId int64, status int) error
}

type NounRepository interface {
	IsWordReal(word string, gameType int) (bool, error)
}
