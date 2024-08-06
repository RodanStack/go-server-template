package apires

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, message string, data interface{}, meta interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  StatusSuccess,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func Created(c *gin.Context, message string, data interface{}, meta interface{}) {
	c.JSON(http.StatusCreated, Response{
		Status:  StatusSuccess,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func NoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

func Error(c *gin.Context, statusCode int, message string, err error, traceID string) {
	c.JSON(statusCode, ErrorResponse{
		Status:  StatusError,
		Message: message,
		Error:   err.Error(),
		TraceID: traceID,
	})
}

func ValidationError(c *gin.Context, message string, fields map[string]string, traceID string) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Status:  StatusError,
		Message: message,
		Fields:  fields,
		TraceID: traceID,
	})
}
