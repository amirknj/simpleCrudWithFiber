package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Id    uint   `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
