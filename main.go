package main

import (
	"gee"
)

func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1>hello gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.String(200, "hello %s", c.Query("name"))
	})
	r.POST("/login", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"name":     c.PostForm("name"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
