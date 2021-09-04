package database

import (
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	db := CreateDatabase()
	item := Item{"title", "desc"}
	db.Add(item)
	if db.Items[0] != item {
		t.Errorf("Item not added")
	}
}

func TestReadAll(t *testing.T) {
	db := CreateDatabase()
	testDb(db)
	read := db.ReadAll()
	if len(read) != len(db.Items) {
		t.Errorf("Read failed")
	}
}

func testDb(db *Database) {
	for i := 0; i < 10; i++ {
		db.Add(Item{
			Title:       "title " + strconv.Itoa(i),
			Description: "desc " + strconv.Itoa(i),
		})
	}
}
