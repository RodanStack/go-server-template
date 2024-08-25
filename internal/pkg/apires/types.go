package apires

type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
)

type Response struct {
	Status  Status      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"` // Additional metadata
}

type ErrorResponse struct {
	Status  Status            `json:"status"`
	Message string            `json:"message"`
	Error   string            `json:"error,omitempty"`
	Fields  map[string]string `json:"fields,omitempty"`   // For validation errors
	TraceID string            `json:"trace_id,omitempty"` // For tracing errors
}
