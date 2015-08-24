package models

import "time"

//Task executed recurrently
type Task struct {
	ID      int    `sql:"type:varchar(100)" gorm:"primary_key" json:"id"`
	Command string `json:"command"`

	CreatedAt time.Time `json:"created_at"`
}
