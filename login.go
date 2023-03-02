package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// model
type User struct {
	gorm.Model
	username string
	password string
	gender   string
	age      int
	status   int
}

// gorm.Model 的定义
// type Model struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
//   }

// main
func main() {

	var DB, err = Connect()
	if err != nil {
		panic("failed to connect to the database")
	}

	// var user User
	// user_test := User{
	// 	id:       8,
	// 	username: "dummy",
	// 	password: "123",
	// 	gender:   "dummy",
	// 	age:      20,
	// 	status:   0,
	// }

	// fmt.Println(user_test)
	DB.AutoMigrate(&User{})
	// result := DB.First(&user)

	result := map[string]interface{}{}
	DB.Model(&User{}).First(&result)

	fmt.Println(result)

	// result := DB.Create(&user_test)
	// fmt.Println(result)

	// fmt.Println(user)
	// user = *Query(DB)
	// fmt.Println(user)

	// Gin_test()
}

// connect and init
func Connect() (gorm.DB, error) {

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

// query for login
func Query(db gorm.DB) *User {
	var user User
	db.Where("username = ?", "Jake").Find(&user)

	return &user
}

// insert record into table
func Insert(db gorm.DB) bool {
	//db.Create()

	return true
}

func Gin_test() {
	r := gin.Default()
	r.GET("hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "helllo",
		})
	})

	r.Run()
}
