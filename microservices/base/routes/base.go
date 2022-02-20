package routes

import (
	"net/http"

	"github.com/JuniorGuerra/boilerplate-go/microservices/base/controllers/base"
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
