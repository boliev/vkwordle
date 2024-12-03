package api

import (
	"fmt"
	cors "github.com/AdhityaRamadhanus/fasthttpcors"
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

	withCors := cors.NewCorsHandler(cors.Options{
		AllowedOrigins:   []string{},
		AllowedHeaders:   []string{},
		AllowedMethods:   []string{},
		AllowCredentials: true,
		AllowMaxAge:      5600,
		Debug:            true,
	})

	api := r.Group("/api/v1", func(rctx *routing.Context) error {
		rctx.Response.Header.Set("Content-Type", "application/json")
		return nil
	})

	api.Post("/game/start", gameHandler.Start)
	api.Post("/game/word", gameHandler.Word)
	api.Get("/game/hint", gameHandler.Hint)
	api.Get("/game/category", gameHandler.Category)

	port := 8180
	fmt.Printf("Serving on localhost:%d\n", port)
	panic(fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), withCors.CorsMiddleware(r.HandleRequest)))
}
