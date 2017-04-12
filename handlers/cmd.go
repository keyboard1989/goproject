package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keyboard1989/goutil"
)

type CmdHandler struct {
	statis *goutil.Statistics
}

func NewCmdHandler(statis *goutil.Statistics) *CmdHandler {
	return &CmdHandler{
		statis: statis,
	}
}

func (h *CmdHandler) Cmd() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "cmd\n")
	}
}
