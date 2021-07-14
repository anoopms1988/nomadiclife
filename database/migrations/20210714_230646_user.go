package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type User_20210714_230646 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20210714_230646{}
	m.Created = "20210714_230646"

	migration.Register("User_20210714_230646", m)
}

// Run the migrations
func (m *User_20210714_230646) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE nomadiclife.`user` ADD verified BOOL DEFAULT false NOT NULL;")
	m.SQL("ALTER TABLE nomadiclife.`user` ADD created DATETIME NULL;")
	m.SQL("ALTER TABLE nomadiclife.`user` ADD updated DATETIME NULL;")
}

// Reverse the migrations
func (m *User_20210714_230646) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
