package models

import "time"

type ToDoTask struct {
	Content string `json:"content"`
}

type Record struct {
	Name      string
	LastTime  time.Time
	ToDoTasks []ToDoTask
}

type RecordCache struct {
	Name           string
	LastTime       time.Time
	MaxLimit       int
	CurrentDailyNo int
}
