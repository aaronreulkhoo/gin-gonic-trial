package database

/* Interfaces for better  */
type Getter interface {
	ReadAll() []Item
}

type Adder interface {
	Add(i Item)
}

/* Database */

type Item struct {
	Title       string
	Description string
}

type Database struct {
	Items []Item
}

func CreateDatabase() *Database {
	return &Database{
		Items: []Item{},
	}
}

func (db *Database) Add(i Item) {
	db.Items = append(db.Items, i)
}

func (db *Database) ReadAll() []Item {
	return db.Items
}
