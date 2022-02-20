package lang

type Message struct {
	ID      string `json:"messageId,omitempty" lang:"langId"`
	Message string `json:"message,omitempty" lang:"langId"`
}
