package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DebugHandler struct {
}

func NewDebugHandler() *DebugHandler {
	return &DebugHandler{}
}

func (h *DebugHandler) Debug() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "debug\n")
	}
}
