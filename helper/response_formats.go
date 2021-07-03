package helper

type Success struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type Failure struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
