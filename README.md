# gin-error-reporting
gin-error-reporting is gin middleware which sends an error to gcp error reporing

## Installation
```
go get -u https://github.com/hiko1129/gin-error-reporting
```

## Usage
```
package main

import (
	"context"
	"os"

	"cloud.google.com/go/errorreporting"
	"github.com/gin-gonic/gin"
	reporting "github.com/hiko1129/gin-error-reporting"
)

func main() {
	ctx := context.Background()
	errorClient, err := errorreporting.NewClient(ctx, os.Getenv("PROJECT_ID"), errorreporting.Config{})

	if err != nil {
		panic(err)
	}

	defer errorClient.Close()

	r := gin.Default()
	r.Use(reporting.New(errorClient)) // set middleware

  r := gin.Default()

  r.GET("/example", func(c *gin.Context) {
    c.Error(errors.New("foo")) // set error
    return
  }

  r.Run()
}
```

gin-error-reporting sends error to gcp error reporting if error is set by `c.Error`
