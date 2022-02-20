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
				Path:    "/:id",
				Handler: handler.GetItem,
			},

			{
				Method:  http.MethodPost,
				Path:    "",
				Handler: handler.NewItem,
			},

			{
				Method:  http.MethodPost,
				Path:    "/form",
				Handler: handler.NewImageItem,
			},

			{
				Method:  http.MethodGet,
				Path:    "/headers",
				Handler: handler.HeadersItem,
			},
		},
	}

	AppRouting = append(AppRouting, routes)
}
