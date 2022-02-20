package lang

type Message struct {
	ID      string `json:"messageId,omitempty" lang:"langId"`
	Message string `json:"message,omitempty" lang:"langId"`
}

type HttpErrors struct {
	GeneralInternalError Message
}

var Errors HttpErrors

func init() {
	Errors = HttpErrors{
		GeneralInternalError: Message{
			ID:      "GENERAL INTERNAL ERROR",
			Message: "General internal error",
		},
	}
}
