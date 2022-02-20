package pipes

type GetItem struct {
	Name string `json:"name"`
}

type NewItem struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      string `json:"age"`
}
