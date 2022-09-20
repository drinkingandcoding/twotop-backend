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

// func (ins Instruction) Value() (driver.Value, error) {
// 	return json.Marshal(ins)
// }

// func (ins *Instruction) Scan(value interface{}) error {
// 	b, ok := value.([]byte)
// 	if !ok {
// 		return errors.New("type assertion to []byte failed")
// 	}
// 	return json.Unmarshal(b, &ins)
// }
