// package main

// import (
// 	"fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// 	"gorm.io/gorm/schema"
// )

// // https://gorm.io/zh_CN/docs/associations.html

// func main() {

// 	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		NamingStrategy: schema.NamingStrategy{
// 			SingularTable: true, //使用单数表名
// 		},
// 		Logger: logger.Default.LogMode(logger.Info),
// 	})
// 	if err != nil {
// 		fmt.Println("数据库连接错误")
// 	}
// 	user := TestUser{Username: "user1", Password: "123456", Status: 0}

// 	db.Create(&user) // 通过数据的指针来创建

// }

// type TestUser struct {
// 	Id       int `gorm:"primaryKey"`
// 	Username string
// 	Password string
// 	Status   int
// }
