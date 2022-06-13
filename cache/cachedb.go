package cache

import (
	"errors"
	"testManabie/models"
	"time"
)

type CacheDb struct {
}

var fakeData map[string]models.RecordCache

func New() *CacheDb {
	fakeData = map[string]models.RecordCache{
		"A": models.RecordCache{
			Name:           "A",
			LastTime:       time.Date(2022, time.March, 2, 11, 12, 30, 30, time.Local),
			MaxLimit:       2,
			CurrentDailyNo: 1,
		},
		"B": models.RecordCache{
			Name:           "B",
			LastTime:       time.Date(2022, time.April, 20, 12, 3, 3, 3, time.Local),
			MaxLimit:       5,
			CurrentDailyNo: 5,
		},
	}
	return &CacheDb{}
}

func (*CacheDb) GetData(username string) (*models.RecordCache, error) {
	val, ok := fakeData[username]
	if ok {
		return &val, nil
	}
	return nil, errors.New("no record")
}
func (*CacheDb) UpdateData(r models.RecordCache) error {
	fakeData[r.Name] = r
	return nil
}

func (*CacheDb) AddCacheRecord(username string, lmt ...int) error {
	maxlmt := 0
	if len(lmt) > 0 {
		maxlmt = lmt[0]
	}

	fakeData[username] = models.RecordCache{
		Name:           username,
		LastTime:       time.Now(),
		MaxLimit:       maxlmt,
		CurrentDailyNo: 0,
	}

	return nil
}
func (*CacheDb) UpdateLastTime(usename string) error {
	rs, ok := fakeData[usename]
	if ok {
		rs.LastTime = time.Now()
		rs.CurrentDailyNo += 1
		fakeData[usename] = rs
		return nil
	}
	return errors.New("failed to update cache")
}
