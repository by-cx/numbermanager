package main

// NumberValue repesents same number as Number but it's used only to set new value into the Number field
type NumberValue struct {
	Number int64 `json:"number"`
}

// Response holds response data
type Response struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Number  int64  `json:"number,omitempty"`
}

// ResponseAlwaysNumber is same as Response, only Number is no omitempty
type ResponseAlwaysNumber struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Number  int64  `json:"number"`
}
