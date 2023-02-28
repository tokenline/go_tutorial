package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 初始化go依赖
// go mod init ginweb

// 安装gin
// go get -u github.com/gin-gonic/gin

func main() {
	Gin4()
}
func Gin4() {
	r := gin.Default()
	r.POST("/fileUpload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "文件为空，请重新上传",
				"success": false,
			})
			return
		}
		// dst:=fmt.Sprint("D:\golerin\demo\ginweb")
		// 绝对路径：D:\golerin\demo\ginweb（完整路径）
		// 相对路径：/fiel/filename.txt
		ctx.SaveUploadedFile(file, "./"+file.Filename)
	})
	r.Run()
}

// Gin3服务（登录）
func Gin3() {
	r := gin.Default()

	r.POST("/user/login", func(c *gin.Context) {
		fmt.Println("是否进来")
		var user3 UserInfo
		c.ShouldBindJSON(&user3)
		fmt.Println("参数：", user3)
		userList := GetUserData()
		if user3.Username == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "用户名不允许为空",
				"success": false,
			})
			return
		}
		user := QueryUser(userList, user3.Username)
		//用户是否存在
		if user == nil {
			fmt.Println("用户不存在，请重新登录")
			c.JSON(http.StatusOK, gin.H{
				"message": "用户不存在，请重新登录",
				"success": false,
			})
			return
		}
		//验证密码是否正确
		if user.Userpass != user3.Userpass {
			fmt.Println("密码错误，请重新登录")
			c.JSON(http.StatusOK, gin.H{
				"message": "密码错误，请重新登录",
				"success": false,
			})
			return
		}
		//验证用户状态
		if user.Status == false {
			fmt.Println("该账号被禁用，无法登录")
			c.JSON(http.StatusOK, gin.H{
				"message": "该账号被禁用，无法登录",
				"success": false,
			})
			return
		}
		fmt.Println("恭喜你，登录成功")

		c.JSON(http.StatusOK, gin.H{
			"message": "恭喜你，登录成功",
			"success": true,
		})
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "welcome",
		// })
	})
	r.Run(":9999") //端口：8080

}

// 查询用户
func QueryUser(userlist []UserData, userName string) *UserData {

	for i := 0; i < len(userlist); i++ {
		if userlist[i].Username == userName {
			return &userlist[i]
		}
	}
	return nil
}

func GetUserData() (userlist []UserData) {

	userlist = append(userlist, UserData{
		Id:       1,
		Username: "user1",
		Userpass: "123456",
		Sex:      "男",
		Age:      13,
		Status:   true,
		Amount:   134,
	})

	userlist = append(userlist, UserData{
		Id:       2,
		Username: "user2",
		Userpass: "112233",
		Sex:      "女",
		Age:      45,
		Status:   true,
		Amount:   456,
	})
	userlist = append(userlist, UserData{
		Id:       3,
		Username: "user3",
		Userpass: "999999",
		Sex:      "男",
		Age:      56,
		Status:   true,
		Amount:   65,
	})

	userlist = append(userlist, UserData{
		Id:       4,
		Username: "user4",
		Userpass: "111111",
		Sex:      "男",
		Age:      34,
		Status:   true,
		Amount:   456,
	})

	userlist = append(userlist, UserData{
		Id:       5,
		Username: "user5",
		Userpass: "222222",
		Sex:      "男",
		Age:      33,
		Status:   true,
		Amount:   33,
	})

	userlist = append(userlist, UserData{
		Id:       6,
		Username: "user6",
		Userpass: "333333",
		Sex:      "女",
		Age:      44,
		Status:   false,
		Amount:   1777,
	})
	return userlist
}

// // | 用户id  | 用户名称  | 密码   | 性别  | 年纪 | 状态 | 用户余额 |
// // | ------ | -------- | ------ | ---- | ---- | ---- | -------- |
// // | 1      | user1    | 123456 | 男   | 13   | 启用 | 134      |
// // | 2      | user2    | 112233 | 女   | 45   | 启用 | 456      |
// // | 3      | user3    | 999999 | 男   | 56   | 启用 | 65       |
// // | 4      | user4    | 111111 | 男   | 34   | 启用 | 456      |
// // | 5      | user5    | 222222 | 男   | 33   | 启用 | 33       |
// // | 6      | user6    | 333333 | 女   | 44   | 禁用 | 1777     |

// // 登录方法如下：
// // 输入用户名、密码
// // 验证用户是否存在
// // 验证密码是否正确
// // 验证状态

// //定义结构体
type UserData struct {
	Id       int    //用户id
	Username string //用户名称
	Userpass string //用户密码
	Sex      string //用户性别
	Age      int    //用户年龄
	Status   bool   //用户状态，0：启用，1：禁用
	Amount   int    //用户余额
}
type UserInfo struct {
	Username string
	Userpass string
}

// 入门
func Gin1() {
	r := gin.Default()

	//Context 包括：request(请求信息)、response（响应信息）
	//路由
	r.GET("hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello go",
		})
	})
	r.Run()
}

// http请求方式、获取参数
func Gin2() {

	r := gin.Default()

	//get：获取
	//post：创建/更新数据
	//put：更新数据
	//delete：删除数据

	//Context 包括：request(请求信息)、response（响应信息）
	//路由
	r.GET("hello/:id", func(c *gin.Context) {
		//Query参数
		name := c.Query("name")
		pwd := c.Query("pwd")
		//路由参数
		id := c.Param("id")
		fmt.Println("name:", name)
		fmt.Println("pwd:", pwd)
		fmt.Println("id:", id)
		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": name,
			"pwd":  pwd,
		})

	})

	r.POST("/add", func(c *gin.Context) {
		//Get、Post：仅支持Query参数
		name := c.Query("name")
		pwd := c.Query("pwd")
		//路由参数
		id := c.Param("id")
		fmt.Println("name:", name)
		fmt.Println("pwd:", pwd)
		fmt.Println("id:", id)
		//Post：body参数、Query参数
		var user UserInfo
		c.ShouldBindJSON(&user)
		c.JSON(http.StatusOK, gin.H{
			"message": "成功", //string
			"success": true, //bool
			"name":    user.Username,
			"pwd":     user.Userpass,
		})
	})
	r.Run()
}

// 路由相关
func router() {

	r := gin.Default()

	//普通路由
	r.GET("hello/:id", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "成功", //string
			"success": true, //bool
		})
	})
	//路由组
	user := r.Group("/user")
	user.GET("/add", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "成功", //string
			"success": true, //bool
		})
	})
	r.Run()
}
