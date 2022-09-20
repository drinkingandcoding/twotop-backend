package database

import (
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	Name        string        `json:"name"`
	Keywords    string        `json:"keywords"`
	Description string        `json:"description"`
	Url         string        `json:"url"`
	Yield       int64         `json:"yield"`
	Ingredients string        `json:"ingredients"`
	Steps       []Instruction `json:"instructions" gorm:"many2many:instruction_steps"`
}

type Instruction struct {
	gorm.Model
	Name string `json:"instructionName"`
	Step string `json:"instructionStep"`
}
