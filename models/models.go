package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Body string `json:"body" gorm:"text;not null;default:null`
	Completed bool `json:"completed" gorm:"default:false`
}

	
