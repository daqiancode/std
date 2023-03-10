package std

// standard result to restful API
type Result struct {
	State       int               `json:"state"` //0: OK, 1: Error
	Data        interface{}       `json:"data,omitempty"`
	Error       string            `json:"error,omitempty"`
	FieldErrors map[string]string `json:"fieldErrors,omitempty"`
	ErrorCode   string            `json:"errorCode,omitempty"`
}
