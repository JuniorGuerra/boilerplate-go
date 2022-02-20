package kernel

import (
	"net/http"

	"github.com/JuniorGuerra/boilerplate-go/kernel/errors"
	"github.com/JuniorGuerra/boilerplate-go/kernel/lang"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
}

func (v Validator) JsonStructure(in interface{}, ctx *gin.Context) (err error) {
	validator := validator.New()
	err = validator.Struct(in)

	if err != nil {
		errors.HttpError(errors.HttpErrors{
			Ctx:    *ctx,
			Status: http.StatusInternalServerError,
			Error: errors.Error{
				Message: lang.Errors.JsonStructureError,
				Extra: errors.Extra{
					"Error": err.Error(),
				},
			},
		})
	}
	return
}
