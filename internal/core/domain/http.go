package domain

type RawJSONMessage []byte

type HTTPResponse struct {
	Data    any    `json:"data"`
	PerPage uint64 `json:"per_page"`
	Page    uint64 `json:"page"`
	Total   uint64 `json:"total"`
}

type HTTPErrorResponse struct {
	Error        RawJSONMessage `json:"error"`
	ErrorMessage string         `json:"errorMessage"`
	Message      string         `json:"message"`
}
