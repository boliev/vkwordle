package game

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// TODO: Add testifylint to golangci-lint
type PuzzleRepositoryMock struct {
	mock.Mock
}

func (m *PuzzleRepositoryMock) GetRandomPuzzle() (*Puzzle, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Puzzle), args.Error(1)
}

type GameRepositoryMock struct {
	mock.Mock
}

func (m *GameRepositoryMock) GetActiveGame(userId int64, gameType int) (*Game, error) {
	args := m.Called(userId, gameType)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Game), args.Error(1)
}

func (m *GameRepositoryMock) AddWord(gameId int64, word string) error {
	args := m.Called(gameId, word)

	return args.Error(0)
}

func (m *GameRepositoryMock) Finish(gameId int64, status int) error {
	args := m.Called(gameId, status)

	return args.Error(0)
}

func (m *GameRepositoryMock) CreateGame(userId int64, word string) (*Game, error) {
	args := m.Called(userId, word)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Game), args.Error(1)
}

func Test_CreateGame_SuccessCreating(t *testing.T) {
	randomPuzzle := &Puzzle{Word: "test"}
	puzzleRepo := new(PuzzleRepositoryMock)
	puzzleRepo.On("GetRandomPuzzle", mock.Anything, mock.Anything).
		Return(randomPuzzle, nil)

	resultGame := &Game{ID: 5}
	gameRepo := new(GameRepositoryMock)
	gameRepo.On("GetActiveGame", mock.Anything, mock.Anything).
		Return(nil, nil)
	gameRepo.On("CreateGame", mock.Anything, mock.Anything).
		Return(resultGame, nil)

	service := NewService(puzzleRepo, gameRepo)
	newGame, err := service.CreateGame(123)
	assert.Nil(t, err)
	assert.NotNil(t, newGame)
	assert.Equal(t, resultGame, newGame)
}

func Test_CreateGame_ActiveGameExists(t *testing.T) {
	puzzleRepo := new(PuzzleRepositoryMock)

	activeGame := &Game{ID: 6}
	gameRepo := new(GameRepositoryMock)
	gameRepo.On("GetActiveGame", mock.Anything, mock.Anything).
		Return(activeGame, nil)

	service := NewService(puzzleRepo, gameRepo)
	newGame, err := service.CreateGame(123)
	assert.Nil(t, err)
	assert.NotNil(t, newGame)
	assert.Equal(t, activeGame, newGame)
}

func Test_CreateGame_Errors(t *testing.T) {
	puzzleRepo := new(PuzzleRepositoryMock)
	gameRepo := new(GameRepositoryMock)

	gameRepo.On("GetActiveGame", mock.Anything, mock.Anything).
		Return(nil, fmt.Errorf("some error"))
	service := NewService(puzzleRepo, gameRepo)
	newGame, err := service.CreateGame(123)
	assert.Nil(t, newGame)
	assert.NotNil(t, err)

	puzzleRepo = new(PuzzleRepositoryMock)
	gameRepo = new(GameRepositoryMock)
	gameRepo.On("GetActiveGame", mock.Anything, mock.Anything).
		Return(nil, nil)
	puzzleRepo.On("GetRandomPuzzle", mock.Anything, mock.Anything).
		Return(nil, fmt.Errorf("some error"))
	service = NewService(puzzleRepo, gameRepo)
	newGame, err = service.CreateGame(123)
	assert.Nil(t, newGame)
	assert.NotNil(t, err)

	puzzleRepo = new(PuzzleRepositoryMock)
	gameRepo = new(GameRepositoryMock)
	randomPuzzle := &Puzzle{Word: "test"}
	puzzleRepo.On("GetRandomPuzzle", mock.Anything, mock.Anything).
		Return(randomPuzzle, nil)
	gameRepo.On("GetActiveGame", mock.Anything, mock.Anything).
		Return(nil, nil)
	gameRepo.On("CreateGame", mock.Anything, mock.Anything).
		Return(nil, fmt.Errorf("some error"))
	service = NewService(puzzleRepo, gameRepo)
	newGame, err = service.CreateGame(123)
	assert.Nil(t, newGame)
	assert.NotNil(t, err)
}

func Test_TryToFinish_In_Progress(t *testing.T) {
	puzzleRepo := new(PuzzleRepositoryMock)
	gameRepo := new(GameRepositoryMock)
	gameRepo.On("Finish", mock.Anything, mock.Anything).Return(nil)

	service := NewService(puzzleRepo, gameRepo)

	game := &Game{
		ID:     1,
		Puzzle: "пират",
		UserId: 123,
		Status: STATUS_IN_PROGRESS,
		Type:   TYPE_5_WORDS,
		Words:  make(map[int]*Word),
	}
	game.Words[0] = &Word{Word: "весна"}
	game.Words[1] = &Word{Word: "пирог"}
	game.Words[2] = &Word{Word: "осень"}
	game.Calc()

	err := service.TryToFinishGame(game)
	assert.Nil(t, err)
	assert.Equal(t, game.Status, STATUS_IN_PROGRESS)
}

func Test_TryToFinish_In_Faild(t *testing.T) {
	puzzleRepo := new(PuzzleRepositoryMock)
	gameRepo := new(GameRepositoryMock)
	gameRepo.On("Finish", mock.Anything, mock.Anything).Return(nil)

	service := NewService(puzzleRepo, gameRepo)

	game := &Game{
		ID:     1,
		Puzzle: "пират",
		UserId: 123,
		Status: STATUS_IN_PROGRESS,
		Type:   TYPE_5_WORDS,
		Words:  make(map[int]*Word),
	}
	game.Words[0] = &Word{Word: "весна"}
	game.Words[1] = &Word{Word: "пирог"}
	game.Words[2] = &Word{Word: "осень"}
	game.Words[3] = &Word{Word: "слово"}
	game.Words[4] = &Word{Word: "аббад"}
	game.Words[5] = &Word{Word: "пирон"}
	game.Calc()

	err := service.TryToFinishGame(game)
	assert.Nil(t, err)
	assert.Equal(t, game.Status, STATUS_FAIL)
}

func Test_TryToFinish_In_Won(t *testing.T) {
	puzzleRepo := new(PuzzleRepositoryMock)
	gameRepo := new(GameRepositoryMock)
	gameRepo.On("Finish", mock.Anything, mock.Anything).Return(nil)

	service := NewService(puzzleRepo, gameRepo)

	game := &Game{
		ID:     1,
		Puzzle: "пират",
		UserId: 123,
		Status: STATUS_IN_PROGRESS,
		Type:   TYPE_5_WORDS,
		Words:  make(map[int]*Word),
	}
	game.Words[0] = &Word{Word: "весна"}
	game.Words[1] = &Word{Word: "пирог"}
	game.Words[2] = &Word{Word: "пират"}
	game.Calc()

	err := service.TryToFinishGame(game)
	assert.Nil(t, err)
	assert.Equal(t, game.Status, STATUS_WON)
}
