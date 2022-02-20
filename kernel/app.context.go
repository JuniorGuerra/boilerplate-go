package kernel

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AppContext struct {
	gin.Context
	Data interface{}
}

func (c AppContext) BindPipe(i interface{}) interface{} {
	err := c.Bind(&i)
	fmt.Println("AAAAAAa")

	if err != nil {
		return err
	}
	// 	if err != nil {
	// 		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
	// 			Message: lang.Message{
	// 				Message: err.Error(),
	// 			},
	// 		})
	// 	}
	// }
	err = c.ShouldBindBodyWith(&i, binding.JSON)
	if err != nil {
		return err
	}

	err = c.ShouldBindBodyWith(&i, binding.XML)
	if err != nil {
		return err
	}

	err = c.ShouldBindHeader(&i)
	if err != nil {
		return err
	}

	err = c.BindQuery(&i)
	if err != nil {
		return err
	}

	err = c.BindWith(&i, binding.Form)
	if err != nil {
		return err
	}

	return i
}
