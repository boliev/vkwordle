package handler

import (
	"encoding/json"
	routing "github.com/qiangxue/fasthttp-routing"
	log "github.com/sirupsen/logrus"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func returnError(rctx *routing.Context, err error, code int, mess string) error {
	log.Errorf("%s", err)
	e := Error{
		Code:    code,
		Message: mess,
	}
	eJson, _ := json.Marshal(e)
	rctx.Response.SetStatusCode(code)
	_, err = rctx.RequestCtx.Write(eJson)
	if err != nil {
		return err
	}

	return nil

	//rctx.RequestCtx.Error(string(eJson), code)
}
