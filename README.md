# CinemaSystem
[![Go](https://github.com/Ltd5552/CinemaSystem/actions/workflows/go.yml/badge.svg)](https://github.com/Ltd5552/CinemaSystem/actions/workflows/go.yml)
## 关于
这是一个Golang操作Mysql数据库的具体实现，学校数据库课程的一次实验，所以偏向于数据库操作     
- 包含增、删、改、查基本操作，应用了连接查询、嵌套查询、分组查询、索引、视图      
- 包含8个实体，9种联系，创建了2个视图，3个索引，购买电影票和发布评论2个事件处理
- 插入、删除操作体现出了关系表的完整性约束，操作、输入异常值时均给出了提示或警告  


## 介绍
应用场景：在线订购电影票及管理系统  
功能简介：  
- 用户可以注册账号，维护个人基本信息，包含昵称、性别、生日、所在地、手机；  
- 用户可以查询场次以及电影信息并购买电影票；  
- 用户可以在购票观影后发布评论，包括电影院评分和电影评分；  
- 用户可以查看电影院信息、放映厅信息；  
- 管理员可以查看并维护场次信息、电影信息；  
- 管理员可以查看取票机剩余可印电影票数；  
- 管理员可以查看评论信息；  
- 系统将记录用户评分并修改电影院和电影均分；  
- 系统将根据售票情况实时修改场次剩余票数；  

## ERD
![CinameERD](/ERD/ERDDiagram1.png)

## 后记
- 实验在验收前三天着手，便没有打算用图形界面，大概半天实现了数据库设计、sql语句及对应函数实现，第二天实现了业务逻辑和原始界面(命令行)，第三天缝缝补补，改了各种bug，晚上勉强验收完成  
- 后续若有机会将结合学习进行web端实现    

_by Ltd 2022.5.21_