package api

import (
	"fmt"
	"github.com/boliev/vkwordle/internal/app/api/handler"
	domain "github.com/boliev/vkwordle/internal/domain/repository/puzzle"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type Api struct {
	puzzleRepo domain.PuzzleRepository
}

func NewApi(puzzleRepo domain.PuzzleRepository) *Api {
	return &Api{
		puzzleRepo: puzzleRepo,
	}
}

func (a *Api) Run() {
	gameHandler := handler.NewGame(a.puzzleRepo)
	r := routing.New()
	api := r.Group("/api/v1")
	api.Get("/start", gameHandler.Start)
	port := 8180
	fmt.Printf("Serving on localhost:%d\n", port)
	panic(fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), r.HandleRequest))
}
