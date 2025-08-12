package main

import (
	"log"
	"net/http"
	"time"

	"JasmineGin/gee"
)

func logger() gee.HandlerFunc {
	return func(c *gee.Context) {
		log.Printf("[%d] %s before,", c.StatusCode, c.Req.RequestURI)
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v,", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func jasmineMiddleware() gee.HandlerFunc {
	return func(c *gee.Context) {
		log.Printf("[%d] %s jasmineMiddleware before ", c.StatusCode, c.Req.RequestURI)
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v jasmineMiddleware", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.Use(logger())

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	jasmine := r.Group("/jasmine")
	jasmine.Use(jasmineMiddleware())
	{
		jasmine.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Jasmine Lee</h1>")
		})

		jasmine.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=jasmine
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/jasmine
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
