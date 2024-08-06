package apires

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandlingMiddleware function to handle errors
//
// # This will activate if you set error to gin context using c.Error(err) where c is gin context
//
// If its expected error you can use apires.Error function to directly send error response
// without setting error to gin context.
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Continue to the next middleware or handler

		if len(c.Errors) > 0 && !c.Writer.Written() { // Check if any errors were recorded and no response has been sent
			err := c.Errors.Last() // Get the last error that occurred
			traceID := c.GetString("trace_id")

			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  StatusError,
				Message: "Internal server error",
				Error:   err.Error(),
				TraceID: traceID,
			})
		}
	}
}
