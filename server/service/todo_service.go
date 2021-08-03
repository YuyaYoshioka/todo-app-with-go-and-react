package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/YuyaYoshioka/todo-app-with-go-and-react/server/db"
	"github.com/YuyaYoshioka/todo-app-with-go-and-react/server/entity"
)

type Service struct{}

type Todo entity.Todo

func (s Service) GetAll() ([]Todo, error) {
	db := db.GetDB()
	var todos []Todo

	if err := db.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (s Service) CreateModel(c *gin.Context) (Todo, error) {
	db := db.GetDB()
	var todo Todo

	if err := c.BindJSON(&todo); err != nil {
		return todo, err
	}

	if err := db.Create(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (s Service) GetByID(id string) (Todo, error) {
	db := db.GetDB()
	var todo Todo

	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (s Service) UpdateByID(id string, c *gin.Context) (Todo, error) {
	db := db.GetDB()
	var todo Todo

	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return todo, err
	}

	if err := c.BindJSON(&todo); err != nil {
		return todo, err
	}

	db.Save(&todo)

	return todo, nil
}

func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var todo Todo

	if err := db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return err
	}

	return nil
}
