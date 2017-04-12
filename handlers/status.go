package handlers

import (
	"net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/keyboard1989/goutil"
)

type StatusHandler struct {
	statis *goutil.Statistics
}

func NewStatusHandler(statis *goutil.Statistics) *StatusHandler {
	return &StatusHandler{
		statis: statis,
	}
}

func (h *StatusHandler) Status() gin.HandlerFunc {
	return func(c *gin.Context) {

		typeStr := c.Request.URL.Query().Get("type")
		subTypeStr := c.Request.URL.Query().Get("sub")

		if typeStr == "" {
			c.String(http.StatusOK, "missing paramater type")
			return
		}

		if subTypeStr == "" {
			res := h.statis.GetInfoByType(typeStr)
			if res == nil {
				c.String(http.StatusOK, "nothing found")
			}

			resJSON, _ := json.Marshal(res)
			c.String(http.StatusOK, string(resJSON))
			return
		}

		res := h.statis.GetInfoByTypeAndSubType(typeStr, subTypeStr)
		if res == nil {
			c.String(http.StatusOK, "nothing found")
			return
		}

		resJSON, _ := json.Marshal(res)
		c.String(http.StatusOK, string(resJSON))
	}
}
