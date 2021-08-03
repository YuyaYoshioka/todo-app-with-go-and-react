package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/YuyaYoshioka/todo-app-with-go-and-react/server/service"
)

type Controller struct {}

func (pc Controller) Index(c *gin.Context) {
	var service todo.Service
	todos, err := service.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, todos)
	}	
}

func (pc Controller) Create(c *gin.Context) {
	var servise todo.Service
	todo, err := servise.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(201, todo)
	}
}

func (pc Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var servise todo.Service
	todo, err := servise.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, todo)
	}
}

func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var servise todo.Service
	todo, err := servise.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(204, todo)
	}
}

func (pc Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var servise todo.Service

	if err := servise.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}