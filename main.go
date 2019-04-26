package main

import (
	golang "github.com/dennysmvn/go-lang/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", golang.CreateTodo)
		v1.GET("/", golang.FetchAllTodo)
		v1.GET("/:id", golang.FetchSingleTodo)
		v1.PUT("/:id", golang.UpdateTodo)
		v1.DELETE("/:id", golang.DeleteTodo)
	}

	router.Run()
}
