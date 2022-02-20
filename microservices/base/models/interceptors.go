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

type Pipe func(ctx kernel.AppContext) interface{}
type Guard func(*gin.Context)

// func(*gin.Context)
type API struct {
	Pipe    Pipe
	Guard   Guard
	Handler func(*gin.Context)
}

func (i *interceptorAPI) ReadData(api API) gin.HandlerFunc {
	return api.Handler
}
