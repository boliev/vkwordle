package handler

import (
	"encoding/json"
	domain "github.com/boliev/vkwordle/internal/domain/repository/puzzle"
	routing "github.com/qiangxue/fasthttp-routing"
)

type StartResponse struct {
	Word string
}

type Game struct {
	puzzleRepo domain.PuzzleRepository
}

func NewGame(puzzleRepo domain.PuzzleRepository) *Game {
	return &Game{
		puzzleRepo: puzzleRepo,
	}
}

func (g *Game) Start(rctx *routing.Context) error {
	puzzle := g.puzzleRepo.GetPuzzle()

	jsn, err := json.Marshal(puzzle)
	if err != nil {
		return err
	}

	_, err = rctx.RequestCtx.Write(jsn)
	if err != nil {
		return err
	}

	return nil
}
