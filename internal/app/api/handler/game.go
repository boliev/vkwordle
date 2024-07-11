package handler

import (
	"encoding/json"
	"github.com/boliev/vkwordle/internal/domain/game"
	routing "github.com/qiangxue/fasthttp-routing"
	log "github.com/sirupsen/logrus"
)

type StartResponse struct {
	Game *game.Game `json:"game"`
}

type Game struct {
	gameService *game.Service
}

func NewGame(gameService *game.Service) *Game {
	return &Game{
		gameService: gameService,
	}
}

func (g *Game) Start(rctx *routing.Context) error {
	newGame, err := g.gameService.CreateGame()
	if err != nil {
		log.Errorf("%s", err)
		rctx.RequestCtx.Error("cannot start game", 400)
	}

	response := StartResponse{
		Game: newGame,
	}

	jsn, err := json.Marshal(response)
	if err != nil {
		return err
	}

	_, err = rctx.RequestCtx.Write(jsn)
	if err != nil {
		return err
	}

	return nil
}
