package reporting

import (
	"log"

	"cloud.google.com/go/errorreporting"
	"github.com/gin-gonic/gin"
)

// New returns a new gin middleware which sends an error to gcp error reporting
func New(client *errorreporting.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()

		if err == nil {
			return
		}

		log.Println(err)

		client.Report(errorreporting.Entry{
			Error: err,
		})
	}
}
