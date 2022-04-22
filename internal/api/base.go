package api

import (
	"github.com/gin-gonic/gin"
	"github.com/willsmil/lover/internal/model"
)

const (
	Success       = 200
	ErrCodeNoAuth = 1001
	ErrCodeParam  = 1002
)

const (
	OK            = "OK"
	ErrMsgNoAuth  = "auth err"
	ErrMsgUnknown = "unknown err"
)

type Handler struct {
	C *gin.Context
}

func Health(ctx *gin.Context) {
	h := Handler{C: ctx}
	h.Response(Success, OK, nil)
}

func (h *Handler) Response(code int, msg string, data interface{}) {
	response := model.Response{
		Code:    code,
		Message: msg,
		Trace:   h.C.GetString("trace"),
		Data:    data,
	}
	h.C.JSON(200, response)
}
