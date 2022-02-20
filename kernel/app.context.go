package kernel

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type AppContext struct {
}

func (c AppContext) BindPipe(ctx gin.Context, i interface{}) error {
	// err := ctx.ShouldBind(&i)

	// if err != nil {
	// 	errors.HttpError(errors.HttpErrors{
	// 		Ctx:    ctx,
	// 		Status: http.StatusInternalServerError,
	// 		Error: errors.Error{
	// 			Message: lang.Errors.GeneralInternalError,
	// 			Extra: errors.Extra{
	// 				"Error": err.Error(),
	// 			},
	// 		},
	// 	})
	// }

	ctx.ShouldBindHeader(&i)
	fmt.Println(i)

	return nil
}
