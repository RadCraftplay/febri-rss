package models

type Message struct {
	OperationName     string  `json:"operation_name"`
	OperationArgsJson *string `json:"args"`
}
