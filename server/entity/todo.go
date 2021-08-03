package entity

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Text string `json:"text"`
	Status bool `json:"status"`
}
