package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) root(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "hello world")
}
