package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keyboard1989/goutil"
)

type ApiHandler struct {
	httpClient *http.Client
	statis     *goutil.Statistics
}

func NewApiHandler(httpClinet *http.Client, statis *goutil.Statistics) *ApiHandler {
	return &ApiHandler{
		httpClient: httpClinet,
		statis:     statis,
	}
}

func (h *ApiHandler) Api() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "hello,world")
	}
}
