package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	index    int
	username string
	password string
	gender   string
	age      int
	status   int
}

func main() {

	var DB, err = connect()
	if err != nil {
		panic("failed to connect to the database")
	}

	var user User

	DB.Where("id = 1").Find(&user)

	fmt.Println(user)
	user = *query(DB)
	fmt.Println(user)

}

func connect() (gorm.DB, error) {

	dsn := "root@tcp(127.0.0.1:3306)/archive?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect to the database")
	}
	fmt.Println("connect successed")
	return *db, nil
}

func query(db gorm.DB) *User {
	var user User
	db.Where("username = ?", "Jake").Find(&user)
	return &user
}
