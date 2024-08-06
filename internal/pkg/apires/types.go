package apires

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"` // Additional metadata
}

type ErrorResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Error   string            `json:"error,omitempty"`
	Fields  map[string]string `json:"fields,omitempty"`   // For validation errors
	TraceID string            `json:"trace_id,omitempty"` // For tracing errors
}

const (
	StatusSuccess = "success"
	StatusError   = "error"
)
