package models

import "time"

var ExecutionPage = 10

//Execution on the system
type Execution struct {
	ID       int    `gorm:"primary_key" json:"id"`
	CPU      int    `json:"cpu"`
	RAM      int    `json:"ram"`
	Bandwith int    `json:"bandwith"`
	Seconds  int    `json:"seconds"`
	Machine  string `json:"machine"`
	TaskID   string `json:"-" validate:"required"`

	CreatedAt time.Time `json:"created_at"`
}
