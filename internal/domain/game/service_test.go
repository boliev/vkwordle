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

func (m *GameRepositoryMock) GetActiveGame(userId int64, gameType int8) (*Game, error) {
	args := m.Called(userId, gameType)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Game), args.Error(1)
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
