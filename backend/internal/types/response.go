package types

type Response struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
