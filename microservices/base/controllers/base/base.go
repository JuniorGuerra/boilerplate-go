package base

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (h Handler) GetItem(ctx *gin.Context) {

	fmt.Println(ctx.Request)
	ctx.JSON(http.StatusOK, "Esto deberia ser un json")
}

func (h Handler) NewItem(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Esto es un new item")
}
