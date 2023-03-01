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


## 开发文档

第一步、最小实现模型

设计：使用Mysql本地数据库，设计简单的用户数据表，使用gorm和gin实现最基本的查询。不使用object存储信息。

数据库设计：第一个最简单的版本中数据库 archive中仅有一个user表

user表的字段包含如下：

 id int 64 主键 非空 非重复 默认自增 // 用户ID

 username varchar 64 非空 非重复 // 用户名

 password varchar 64 非空 //密码

 gender varchar 20 非空 //在后续版本中做约束

 age int 20 非空 //年龄

 状态 int 8 default 0 // 状态 0默认启用 在后续版本中做约束


### 3月1日 

初步写了connect()和query()两个函数用于初始化数据库和查询，已经连接到本地Mysql但是返回全部都是0和空字符串

怀疑是model里的id写成了index导致没读出来，但是改成id还是返回0和空字符串。

目前不清楚原因。



