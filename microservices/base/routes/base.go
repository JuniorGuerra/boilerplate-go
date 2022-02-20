package routes

import (
	"fmt"
	"net/http"

	"github.com/JuniorGuerra/boilerplate-go/microservices/base/controllers/base"
	"github.com/JuniorGuerra/boilerplate-go/microservices/base/models"
	"github.com/gin-gonic/gin"
)

func init() {
	const subjet = "base"
	handler := base.Handler{}

	routes := Group{
		Prefix: subjet,
		Routes: []Route{
			{
				Method:  http.MethodGet,
				Path:    "",
				Handler: handler.GetItem,
				// Esto aun no funciona
				Data: models.API{
					Pipe: func(i interface{}, c *gin.Context) {
						fmt.Println("AAAAAa")
					},
				},
			},
			{
				Method:  http.MethodPost,
				Path:    "",
				Handler: handler.NewItem,
			},
		},
	}

	AppRouting = append(AppRouting, routes)
}
