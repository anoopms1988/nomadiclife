package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type User_20210628_221252 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20210628_221252{}
	m.Created = "20210628_221252"

	migration.Register("User_20210628_221252", m)
}

// Run the migrations
func (m *User_20210628_221252) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE nomadiclife.`user` MODIFY COLUMN username varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL;")
}

// Reverse the migrations
func (m *User_20210628_221252) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
