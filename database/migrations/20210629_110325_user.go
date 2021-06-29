package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type User_20210629_110325 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20210629_110325{}
	m.Created = "20210629_110325"

	migration.Register("User_20210629_110325", m)
}

// Run the migrations
func (m *User_20210629_110325) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	//m.SQL("ALTER TABLE nomadiclife.`user` MODIFY COLUMN username varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL;")
	m.SQL("ALTER TABLE nomadiclife.`user` ADD CONSTRAINT email_constraint UNIQUE (email)")
	m.SQL("ALTER TABLE nomadiclife.`user` ADD CONSTRAINT username_constraint UNIQUE (username)")
}

// Reverse the migrations
func (m *User_20210629_110325) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
