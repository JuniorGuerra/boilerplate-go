package base

import (
	"fmt"
	"net/http"

	"github.com/JuniorGuerra/boilerplate-go/kernel/errors"
	"github.com/JuniorGuerra/boilerplate-go/kernel/lang"
	"github.com/JuniorGuerra/boilerplate-go/microservices/base/pipes"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (h Handler) GetItem(ctx *gin.Context) {
	// var appCtx kernel.AppContext

	fmt.Println(ctx.Request.Body)

	ctx.JSON(http.StatusOK, "Esto deberia ser un json")
}

func (h Handler) NewItem(ctx *gin.Context) {
	var pipe pipes.NewItem
	err := ctx.ShouldBindJSON(&pipe)

	if err != nil {
		errors.HttpError(errors.HttpErrors{
			Ctx:    *ctx,
			Status: http.StatusInternalServerError,
			Error: errors.Error{
				Message: lang.Errors.GeneralInternalError,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, pipe)
}
