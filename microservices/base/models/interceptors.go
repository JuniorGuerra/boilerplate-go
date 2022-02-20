package models

import (
	"github.com/JuniorGuerra/boilerplate-go/kernel"
	"github.com/gin-gonic/gin"
)

type interceptorAPI struct {
	i kernel.Interceptor
}

var InterceptorAPI = interceptorAPI{}

func init() {
	InterceptorAPI = interceptorAPI{
		kernel.Interceptor{},
	}
}

type Pipe func(interface{}, *gin.Context)
type Guard func(*gin.Context)

type API struct {
	Pipe  Pipe
	Guard Guard
}

func ReadData(kernel.AppContext) interface{} {
	return nil
}
