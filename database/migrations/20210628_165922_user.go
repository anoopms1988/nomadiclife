package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type User_20210628_165922 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20210628_165922{}
	m.Created = "20210628_165922"

	migration.Register("User_20210628_165922", m)
}

// Run the migrations
func (m *User_20210628_165922) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(" ALTER TABLE nomadiclife.`user` ADD address TEXT NULL;")
	m.SQL("ALTER TABLE nomadiclife.`user` ADD phone_number varchar(100) NULL;")
	m.SQL("ALTER TABLE nomadiclife.`user` ADD password varchar(100) NOT NULL;")
	m.SQL("ALTER TABLE nomadiclife.`user` ADD username varchar(100) NOT NULL;")
}

// Reverse the migrations
func (m *User_20210628_165922) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
