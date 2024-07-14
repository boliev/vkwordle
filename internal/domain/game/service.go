package game

type Service struct {
	puzzleRepo PuzzleRepository
	gameRepo   GameRepository
}

func NewService(puzzleRepo PuzzleRepository, gameRepo GameRepository) *Service {
	return &Service{
		puzzleRepo: puzzleRepo,
		gameRepo:   gameRepo,
	}
}

func (s *Service) CreateGame(userId int64) (*Game, error) {

	activeGame, err := s.gameRepo.GetActiveGame(userId, TYPE_5_WORDS)
	if err != nil {
		return nil, err
	}

	if activeGame != nil {
		return activeGame, nil
	}

	puzzle, err := s.puzzleRepo.GetRandomPuzzle()
	if err != nil {
		return nil, err
	}

	game, err := s.gameRepo.CreateGame(userId, puzzle.Word)
	if err != nil {
		return nil, err
	}

	return game, nil
}
