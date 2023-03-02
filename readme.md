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


### 3月2日

预计加入gin框架 写一个最简单的api接口

简单的api跑完了没啥问题 先尝试解决gorm的这个问题再把两个写到一起

应该就能实现预定的/auth/login接口

gorm问题解决，原因是因为在定义model的时候并没有定义它是一个gorm model

调gorm的时候第一次接触到软删除和硬删除在DB.First(&user)

问题：

1.通过DB.First获取的这个对象，存到map中是否合适？网上参考案例存的是map[string]中的
这显然不合适，原定User这个结构体中既有int，也有string.

2.直接通过 result := DB.First获取这个东西也行，但是打印出来的这个
result里的东西好像是个地址.它究竟是什么？