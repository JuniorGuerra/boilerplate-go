package pipes

type GetItem struct {
	Name string `json:"name"`
}

type NewItem struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Age      int    `json:"age" validate:"required"`
}
