// package main

// import (
// 	"fmt"
// )

// // 需求05
// // | 用户id  | 用户名称  | 密码   | 性别  | 年纪 | 状态 | 用户余额 |
// // | ------ | -------- | ------ | ---- | ---- | ---- | -------- |
// // | 1      | user1    | 123456 | 男   | 13   | 启用 | 134      |
// // | 2      | user2    | 112233 | 女   | 45   | 启用 | 456      |
// // | 3      | user3    | 999999 | 男   | 56   | 启用 | 65       |
// // | 4      | user4    | 111111 | 男   | 34   | 启用 | 456      |
// // | 5      | user5    | 222222 | 男   | 33   | 启用 | 33       |
// // | 6      | user6    | 333333 | 女   | 44   | 禁用 | 1777     |
// // 1、一共有多少个用户
// // 2、所有用户余额是多少
// // 3、打印余额最多的用户
// // 4、打印余额最少的用户
// // 5、打印年纪最大的用户
// // 6、打印年纪最小的用户
// // 7、打印户余额超过100的用户
// // 8、打印年龄大于或者等于25的用户
// // 9、打印状态为启用的用户
// // 10、登录方法如下：
// // 输入用户名、密码
// // 验证用户是否存在
// // 验证密码是否正确
// // 验证状态

// type User struct {
// 	index    int
// 	username string
// 	password string
// 	gender   string
// 	age      int
// 	status   bool
// 	balance  int
// }

// // func main() {

// // 	var userList []user
// // 	userList = append(userList, user{
// // 		index:    1,
// // 		username: "user1",
// // 		password: "123456",
// // 		gender:   "male",
// // 		age:      13,
// // 		status:   true,
// // 		balance:  134,
// // 	})

// // 	userList = append(userList, user{
// // 		index:    2,
// // 		username: "user2",
// // 		password: "112234",
// // 		gender:   "male",
// // 		age:      45,
// // 		status:   true,
// // 		balance:  456,
// // 	})

// // 	userList = append(userList, user{
// // 		index:    3,
// // 		username: "user2",
// // 		password: "112234",
// // 		gender:   "male",
// // 		age:      45,
// // 		status:   true,
// // 		balance:  90,
// // 	})
// // 	//fmt.Println(userList)
// // 	//fmt.Println(len(userList))

// // 	// 2、所有用户余额是多少
// // 	var totalBalance = 0
// // 	for i := 0; i < len(userList); i++ {
// // 		totalBalance += userList[i].balance
// // 	}
// // 	//fmt.Println("所有用户余额", totalBalance)

// // 	// 3、打印余额最多的用户
// // 	var maxBalance int
// // 	var maxBalanceUserName string
// // 	for i := 0; i < len(userList); i++ {
// // 		if userList[i].balance > maxBalance {
// // 			maxBalance = userList[i].balance
// // 			maxBalanceUserName = userList[i].username
// // 		}
// // 		// balance = append(balance, userList[i].balance)
// // 		// sort.Ints(balance)
// // 	}
// // 	fmt.Println("余额最多的用户", maxBalanceUserName)

// // 	// 4、打印余额最少的用户
// // 	var minBalance int
// // 	var minBalanceUserName string
// // 	for i := 0; i < len(userList); i++ {
// // 		if i == 0 {
// // 			//第一次（初始化）
// // 			minBalance = userList[i].balance
// // 			minBalanceUserName = userList[i].username
// // 		}
// // 		if userList[i].balance < minBalance {
// // 			minBalance = userList[i].balance
// // 			minBalanceUserName = userList[i].username
// // 		}
// // 	}
// // 	fmt.Println("余额最多的用户", minBalanceUserName)

// // 	// 7、打印户余额超过100的用户
// // 	for i := 0; i < len(userList); i++ {
// // 		if userList[i].balance > 100 {
// // 			fmt.Println("余额超过100的用户", userList[i].username)
// // 		}
// // 	}
// // }

// // 10、登录方法如下：
// // 输入用户名、密码
// // 验证用户是否存在
// // 验证密码是否正确
// // 验证状态

// func main() {

// 	var userList []user
// 	userList = append(userList, user{
// 		index:    1,
// 		username: "user1",
// 		password: "123456",
// 		gender:   "male",
// 		age:      13,
// 		status:   true,
// 		balance:  134,
// 	})

// 	userList = append(userList, user{
// 		index:    2,
// 		username: "user2",
// 		password: "112234",
// 		gender:   "male",
// 		age:      45,
// 		status:   true,
// 		balance:  456,
// 	})

// 	userList = append(userList, user{
// 		index:    3,
// 		username: "user2",
// 		password: "112234",
// 		gender:   "male",
// 		age:      45,
// 		status:   false,
// 		balance:  90,
// 	})

// 	fmt.Print("请输入用户名：")
// 	var username string
// 	fmt.Scanln(&username)
// 	fmt.Print("请输入密码：")
// 	var password string
// 	fmt.Scanln(&password)
// 	var flag bool = authorize(username, password, userList)
// 	if flag == true {
// 		fmt.Println("用户名密码正确")

// 	} else {
// 		fmt.Println("用户名密码错误")
// 	}

// }

// func authorize(username string, password string, userList []user) bool {
// 	for i := 0; i < len(userList); i++ {
// 		if userList[i].username == username && userList[i].password == password {
// 			return true
// 		}
// 	}
// 	return false
// }
