package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type User_20210628_144749 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20210628_144749{}
	m.Created = "20210628_144749"

	migration.Register("User_20210628_144749", m)
}

// Run the migrations
func (m *User_20210628_144749) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE user ADD email VARCHAR(255)")
}

// Reverse the migrations
func (m *User_20210628_144749) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
