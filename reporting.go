package reporting

import (
	"cloud.google.com/go/errorreporting"
	"github.com/gin-gonic/gin"
)

// New returns a new gin middleware which send error report to gcp error reporting
func New(client *errorreporting.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := c.Errors.Last()

		client.Report(errorreporting.Entry{
			Error: err,
		})
	}
}
