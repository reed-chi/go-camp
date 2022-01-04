package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func Healthz(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{"message": "OK"},
	)
	return
}

func SetHeader(c *gin.Context) {
	//get request header
	geek := c.GetHeader("Geek")
	//set header to response
	c.Header("Country", geek)
}

func GetEnv(c *gin.Context) {
	version := os.Getenv("VERSION")
	c.Header("Version", version)
}

func main() {
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.GET("/header", SetHeader)
	r.HEAD("/header", SetHeader)
	r.GET("/healthz", Healthz)
	r.GET("/version", GetEnv)
	r.HEAD("/version", GetEnv)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
