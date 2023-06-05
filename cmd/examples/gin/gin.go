package main

import (
	"fmt"
	example "github.com/micklove/tracr"
	mid "github.com/micklove/tracr/internal/middleware"
	"github.com/micklove/tracr/internal/tracr"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// TODO - Add middleware

func main() {
	// Create the correlation ID options, in this case, we will simply hard code in the funcs that will
	// retrieve the correlation ID http header name and the correlation ID value. This allows the caller to
	// use whatever strategy they want, to generate the correlation ID e.g. DB Sequence, UUID, GUID, etc... and / or
	// to use their preferred strategy for naming the correlation id http header (e.g. from env var, config , ssm, etc..)
	cidHttpHeaderName := "my-trace-header"
	correlationIDOptions := tracr.CorrelationIDOptions{
		CorrelationIDHttpHeaderFn: func() (string, error) { return cidHttpHeaderName, nil },
		CorrelationIDGeneratorFn:  func() (string, error) { return "b63a65cc-20fa-4b17-97ad-b796bdb6d338", nil },
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	r := gin.Default()
	r.Use(mid.MiddlewareCorrelationIDGin(correlationIDOptions, nil))

	r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	r.GET("/:name", func(c *gin.Context) {
		// get my-trace-header from the context
		cid, err := tracr.GetCID(c.Request.Context())
		if err != nil {
			log.Printf("GetCID() = returned error %v", err)
		}

		// get the my-trace-header from the request
		headerCid := c.Request.Header.Get(cidHttpHeaderName)

		log.Printf("Header %s = %s", cidHttpHeaderName, headerCid)
		log.Printf("Context CID = %s", cid)

		// echo the correlation id header and value in the response
		c.Header(cidHttpHeaderName, cid)

		s := example.NewService()

		cidFromService, err := s.Echo(c.Request.Context())
		if err != nil {
			log.Printf("Echo() = returned error %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"incoming_request_correlation_id": fmt.Sprintf("%s", headerCid),
			"context_cid_from_service":        fmt.Sprintf("%s", cidFromService),
		})
	})

	r.Run(fmt.Sprintf(":%s", port))
}
