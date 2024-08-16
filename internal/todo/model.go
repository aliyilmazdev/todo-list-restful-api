package todo

import "time"

type Todo struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement:true;not null;unique"`
	Title       string    `json:"title"`
	Description string 	  `json:"description"`
	Completed   bool   	  `json:"completed"`
	Time 	  	time.Time `json:"time"`
}