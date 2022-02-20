package pipes

import "mime/multipart"

type GetItem struct {
	Id string `param:"id"`
}

type NewItem struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

type NewImageItem struct {
	Name   string                `form:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

type HeadersItem struct {
	Rate   *int    `header:"Rate"`
	Domain *string `header:"Domain"`
}
