package main

import (
	"errors"
	"log"
	"net/http"
	"testManabie/models"
	"testManabie/utils"
	"time"
)

type RequestPayload struct {
	Username string          `json:"username"`
	TodoTask models.ToDoTask `json:"todotask"`
}

func (app *Config) InsertTask(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	_ = utils.ReadJSON(w, r, &requestPayload)

	ck, err := app.CheckCachedDb(requestPayload.Username)
	if err != nil {
		utils.ErrorJSON(w, errors.New("cannot check current status"))
		return
	}

	if !ck {
		utils.ErrorJSON(w, errors.New("above limit"))
		return
	}

	err = app.Db.AddData(requestPayload.Username, requestPayload.TodoTask)

	if err != nil {
		utils.ErrorJSON(w, errors.New("failed to add task"))
		return
	}

	err = app.CacheDb.UpdateLastTime(requestPayload.Username)

	if err != nil {
		utils.ErrorJSON(w, errors.New("failed to update cache"))
		return
	}

	utils.WriteJSON(w, http.StatusAccepted, utils.JsonResponse{
		Error:   false,
		Message: "add todotask successful",
	})
}

func (app *Config) CheckCachedDb(username string) (bool, error) {
	r, err := app.CacheDb.GetData(username)

	if err != nil {
		err = app.CacheDb.AddCacheRecord(username, 1)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	log.Println(r.Name, r.CurrentDailyNo, r.MaxLimit, r.LastTime)

	tn := time.Now()
	td := time.Date(tn.Year(), tn.Month(), tn.Day(), 0, 0, 0, 0, time.Local)

	//Reset count
	if r.LastTime.Before(td) {
		r.CurrentDailyNo = 0
		app.CacheDb.UpdateData(*r)
	}

	log.Println(r.CurrentDailyNo)

	return r.CurrentDailyNo < r.MaxLimit, nil
}

func (app *Config) AddToDatabase(username string, tdt models.ToDoTask) error {
	err := app.Db.AddData(username, tdt)
	return err
}
