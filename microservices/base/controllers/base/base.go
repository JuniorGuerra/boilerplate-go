package base

import (
	"net/http"

	"github.com/JuniorGuerra/boilerplate-go/kernel"
	"github.com/JuniorGuerra/boilerplate-go/microservices/base/pipes"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (h Handler) GetItem(ctx *gin.Context) {
	var pipe pipes.GetItem

	err := kernel.AppContext{}.BindPipe(*ctx, &pipe)

	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, pipe)
}

func (h Handler) NewItem(ctx *gin.Context) {
	var pipe pipes.NewItem

	err := kernel.AppContext{}.BindPipe(*ctx, &pipe)

	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, pipe)
}

func (h Handler) NewImageItem(ctx *gin.Context) {
	var pipe pipes.NewImageItem

	err := kernel.AppContext{}.BindPipe(*ctx, &pipe)

	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, pipe.Name)
}

func (h Handler) HeadersItem(ctx *gin.Context) {
	var pipe pipes.HeadersItem

	err := kernel.AppContext{}.BindPipe(*ctx, pipe)

	if err != nil {
		return
	}

	if err = ctx.ShouldBindHeader(&pipe); err != nil {
		ctx.JSON(200, err)
	}

	ctx.JSON(http.StatusOK, pipe)
}
