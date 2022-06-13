package main

import (
	"testManabie/cache"
	"testManabie/db"
	"testing"
)

func TestCacheDbGetData(t *testing.T) {
	cachedb := cache.New()
	r, err := cachedb.GetData("A")

	if err != nil {
		t.Errorf("failed to get data")
	}

	if r.Name != "A" {
		t.Errorf("failed to get A")
	}
}

func TestDatabaseGetData(t *testing.T) {
	db := db.New()

	r := db.GetData("A")

	if r.Name != "A" {
		t.Errorf("failed to get A")
	}
}
