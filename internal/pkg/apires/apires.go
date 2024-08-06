package apires

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success function to send success response.
func Success(c *gin.Context, message string, data interface{}) {
	if !c.Writer.Written() {
		c.JSON(http.StatusOK, Response{
			Status:  StatusSuccess,
			Message: message,
			Data:    data,
		})
	}
}

// Created function to send created response.
func Created(c *gin.Context, message string, data interface{}) {
	if !c.Writer.Written() {
		c.JSON(http.StatusCreated, Response{
			Status:  StatusSuccess,
			Message: message,
			Data:    data,
		})
	}
}

// NoContent function to send no content response.
func NoContent(c *gin.Context) {
	if !c.Writer.Written() {
		c.JSON(http.StatusNoContent, nil)
	}
}

// Error function to send error response.
func Error(c *gin.Context, statusCode int, message string, err error, traceID string) {
	if !c.Writer.Written() {
		c.JSON(statusCode, ErrorResponse{
			Status:  StatusError,
			Message: message,
			Error:   err.Error(),
			TraceID: traceID,
		})
	}
}

// ValidationError function to send validation error response.
func ValidationError(c *gin.Context, message string, fields map[string]string, traceID string) {
	if !c.Writer.Written() {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  StatusError,
			Message: message,
			Fields:  fields,
			TraceID: traceID,
		})
	}
}
