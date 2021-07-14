package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type User_20210713_191952 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20210713_191952{}
	m.Created = "20210713_191952"

	migration.Register("User_20210713_191952", m)
}

// Run the migrations
func (m *User_20210713_191952) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *User_20210713_191952) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
