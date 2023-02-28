# 项目文档

1、Go语法
2、Gin+Gorm自我练习
3、Gin+Gorm登录
用户登录方法
（1）输入用户名、密码
（2）验证用户是否存在
（3）验证密码是否正确
（4）验证用户状态
通过Gin，结合Gorm做一个登录
接口文档
接口地址：/auth/login
请求方式：POST
Body参数：{
    "username":"user1",
    "password":"123456"
}

