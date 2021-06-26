package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type User_20210627_000210 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20210627_000210{}
	m.Created = "20210627_000210"

	migration.Register("User_20210627_000210", m)
}

// Run the migrations
func (m *User_20210627_000210) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE user(id INT AUTO_INCREMENT PRIMARY KEY, first_name VARCHAR(255), last_name VARCHAR(255))")
}

// Reverse the migrations
func (m *User_20210627_000210) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
