package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type User_20210714_232400 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20210714_232400{}
	m.Created = "20210714_232400"

	migration.Register("User_20210714_232400", m)
}

// Run the migrations
func (m *User_20210714_232400) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE user ADD verification_token VARCHAR(255)")
}

// Reverse the migrations
func (m *User_20210714_232400) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
