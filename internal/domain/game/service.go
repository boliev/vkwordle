package game

import "fmt"

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

	activeGame, err := s.GetActiveGame(userId)
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

func (s *Service) GetActiveGame(userId int64) (*Game, error) {
	activeGame, err := s.gameRepo.GetActiveGame(userId, TYPE_5_WORDS)
	if err != nil {
		return nil, err
	}

	return activeGame, nil
}

func (s *Service) AddWord(game *Game, word string) error {
	if game.Type == TYPE_5_WORDS && len([]rune(word)) != 5 {
		return fmt.Errorf("invalid word %s", word)
	}

	err := s.gameRepo.AddWord(game.ID, word)
	if err != nil {
		return err
	}
	game.AddWord(word)
	game.Calc()

	return nil
}
