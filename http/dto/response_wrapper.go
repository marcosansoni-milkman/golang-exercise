package dto

type Response struct {
	Success bool          `json:"success"`
	Result  ResultWrapper `json:"result,omitempty"`
	Errors  []Error       `json:"errors,omitempty"`
}

type ResultWrapper struct {
	Message string `json:"message"`
	Entity  Entity `json:"entity,omitempty"`
}

type Entity struct {
	User User `json:"user"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
