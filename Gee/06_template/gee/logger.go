package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Content) {
		// Start time
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s is in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
