package game

import (
	"fmt"
	"sync"
)

type Service struct {
	puzzleRepo PuzzleRepository
	gameRepo   GameRepository
	nounRepo   NounRepository
	mu         sync.Mutex
}

func NewService(puzzleRepo PuzzleRepository, gameRepo GameRepository, nounRepo NounRepository) *Service {
	return &Service{
		puzzleRepo: puzzleRepo,
		gameRepo:   gameRepo,
		nounRepo:   nounRepo,
		mu:         sync.Mutex{},
	}
}

func (s *Service) CreateGame(userId int64) (*Game, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

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

	isReal, err := s.nounRepo.IsWordReal(word, game.Type)
	if err != nil {
		return err
	}

	if !isReal {
		return fmt.Errorf("there is no word %s", word)
	}

	err = s.gameRepo.AddWord(game.ID, word)
	if err != nil {
		return err
	}
	game.AddWord(word)
	game.Calc()
	err = s.TryToFinishGame(game)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) TryToFinishGame(game *Game) error {
	if game.IsLastWordCorrect() {
		game.Status = STATUS_WON
		err := s.gameRepo.Finish(game.ID, STATUS_WON)
		if err != nil {
			return err
		}
	} else if game.Type == TYPE_5_WORDS && game.WordsLen() == 6 {
		game.Status = STATUS_FAIL
		err := s.gameRepo.Finish(game.ID, STATUS_FAIL)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) GetPuzzle(word string) (*Puzzle, error) {
	return s.puzzleRepo.GetPuzzle(word)
}
