package handler

import (
	"encoding/json"
	"fmt"
	"github.com/boliev/vkwordle/internal/domain/game"
	routing "github.com/qiangxue/fasthttp-routing"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type GameResponse struct {
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
	userId, err := getUserId(rctx)
	if err != nil {
		log.Errorf("%s", err)
		rctx.RequestCtx.Error("cannot start game. User not found", 400)
		return nil
	}

	newGame, err := g.gameService.CreateGame(userId)
	if err != nil {
		log.Errorf("%s", err)
		rctx.RequestCtx.Error("cannot start game", 400)
		return nil
	}

	response := GameResponse{
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

func (g *Game) Word(rctx *routing.Context) error {
	userId, err := getUserId(rctx)
	if err != nil {
		log.Errorf("%s", err)
		rctx.RequestCtx.Error("user not found", 400)
		return nil
	}

	type postRequest struct {
		Word string `json:"word"`
	}
	request := &postRequest{}
	err = json.Unmarshal(rctx.RequestCtx.PostBody(), request)
	if err != nil {
		log.Errorf("%s", err)
		rctx.RequestCtx.Error("invalid request", 400)
		return nil
	}

	activeGame, err := g.gameService.GetActiveGame(userId)
	if err != nil {
		log.Errorf("%s", err)
		rctx.RequestCtx.Error("active game not found", 400)
		return nil
	}

	if activeGame == nil {
		log.Errorf("active game not found for user %d", userId)
		rctx.RequestCtx.Error("active game not found", 400)
		return nil
	}

	if !activeGame.InProgress() {
		log.Errorf("game is finished %d", activeGame.ID)
		rctx.RequestCtx.Error("game is finished. You should start a new game", 400)
		return nil
	}

	err = g.gameService.AddWord(activeGame, request.Word)

	if err != nil {
		log.Errorf("%s", err)
		rctx.RequestCtx.Error("cannot add word to the game", 400)
		return nil
	}

	response := GameResponse{
		Game: activeGame,
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

func (g *Game) Hint(rctx *routing.Context) error {
	puzzle, err := g.getUserPuzzle(rctx)
	if err != nil {
		log.Errorf("%s", err)
		rctx.RequestCtx.Error("hint for active game not found", 400)
		return nil
	}

	type hintResponse struct {
		Hint string `json:"hint"`
	}

	response := hintResponse{
		Hint: puzzle.Hint,
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

func (g *Game) Category(rctx *routing.Context) error {
	puzzle, err := g.getUserPuzzle(rctx)
	if err != nil {
		log.Errorf("%s", err)
		rctx.RequestCtx.Error("hint for active game not found", 400)
		return nil
	}

	type hintResponse struct {
		Category string `json:"category"`
	}

	response := hintResponse{
		Category: puzzle.Category,
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

func (g *Game) getUserPuzzle(rctx *routing.Context) (*game.Puzzle, error) {
	userId, err := getUserId(rctx)
	if err != nil {
		return nil, err
	}

	activeGame, err := g.gameService.GetActiveGame(userId)
	if err != nil || activeGame == nil {
		return nil, err
	}

	puzzle, err := g.gameService.GetPuzzle(activeGame.Puzzle)
	if err != nil {
		return nil, err
	}

	return puzzle, nil
}

func getUserId(rctx *routing.Context) (int64, error) {
	userIdBytes := rctx.Request.Header.Peek("userId")
	if len(userIdBytes) == 0 {
		return 0, fmt.Errorf("user not found")
	}
	userId, err := strconv.Atoi(string(userIdBytes))
	if err != nil {
		return 0, err
	}

	return int64(userId), nil
}
