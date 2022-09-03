package main

import (
	"gee_web/gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// start timer
		t := time.Now()
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
		c.Next()
		log.Printf("-----22222")
	}
}

func onlyForV1() gee.HandlerFunc {
	return func(c *gee.Context) {
		// start timer
		t := time.Now()
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
		c.Next()
		log.Printf("-----11111")
	}
}

func main() {
	r := gee.New()

	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	v1.Use(onlyForV2())
	{
		v1.Use(onlyForV1())
		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s\n", c.Path)
		})
	}

	r.POST("/login", func(c *gee.Context) {
		// 业务代码
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
