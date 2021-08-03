package server

import (
	"github.com/YuyaYoshioka/todo-app-with-go-and-react/server/controller"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := router()
	router.Run()
}

func router() *gin.Engine {
	router := gin.Default()
	
	t := router.Group("/todos")
	{
		ctrl := todo.Controller{}
		t.GET("", ctrl.Index)
		t.GET("/:id", ctrl.Show)
		t.POST("", ctrl.Create)
		t.PUT("/:id", ctrl.Update)
		t.DELETE("/:id", ctrl.Delete)
	}

	return router
}
