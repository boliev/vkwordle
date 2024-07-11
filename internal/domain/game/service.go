package game

type Service struct {
	puzzleRepo PuzzleRepository
}

func NewService(PuzzleRepo PuzzleRepository) *Service {
	return &Service{
		puzzleRepo: PuzzleRepo,
	}
}

func (g *Service) CreateGame() (*Game, error) {
	puzzle, err := g.puzzleRepo.GetPuzzle()
	if err != nil {
		return nil, err
	}

	return &Game{
		ID:     1,
		Status: puzzle.Word,
		Words:  make(map[uint]Word),
	}, nil
}
