package reporting

import (
	"log"

	"cloud.google.com/go/errorreporting"
	"github.com/gin-gonic/gin"
)

// New returns a new gin middleware which send error report to gcp error reporting
func New(client *errorreporting.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		log.Println(err)

		client.Report(errorreporting.Entry{
			Error: err,
		})
	}
}
