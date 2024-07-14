package api

import (
	"fmt"
	"github.com/boliev/vkwordle/internal/app/api/handler"
	"github.com/boliev/vkwordle/internal/domain/game"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type Api struct {
	gameService *game.Service
}

func NewApi(gameService *game.Service) *Api {
	return &Api{
		gameService: gameService,
	}
}

func (a *Api) Run() {
	gameHandler := handler.NewGame(a.gameService)
	r := routing.New()

	api := r.Group("/api/v1", func(rctx *routing.Context) error {
		rctx.Response.Header.Set("Content-Type", "application/json")
		return nil
	})

	api.Get("/game/start", gameHandler.Start)
	port := 8180
	fmt.Printf("Serving on localhost:%d\n", port)
	panic(fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), r.HandleRequest))
}
