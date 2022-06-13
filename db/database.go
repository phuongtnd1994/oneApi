package db

import (
	"testManabie/models"
	"time"
)

type Db struct{}

var fakeData map[string]models.Record

func New() *Db {
	fakeData = map[string]models.Record{
		"A": models.Record{
			Name:     "A",
			LastTime: time.Date(2022, time.January, 10, 12, 20, 30, 30, time.Local),
			ToDoTasks: []models.ToDoTask{
				models.ToDoTask{
					Content: "ABC",
				},
			},
		},
	}
	return &Db{}
}

func (*Db) GetData(username string) models.Record {
	return fakeData[username]
}

func (*Db) AddData(username string, tdt models.ToDoTask) error {
	if rc, ok := fakeData[username]; ok {
		rc.ToDoTasks = append(rc.ToDoTasks, tdt)
		fakeData[username] = rc
		return nil
	} else {
		fakeData[username] = models.Record{
			Name:     username,
			LastTime: time.Now(),
			ToDoTasks: []models.ToDoTask{
				tdt,
			},
		}

		return nil
	}
}
