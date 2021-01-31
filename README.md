# api-go

以 Gin，Gorm 为基础的 Golang WebAPI 项目开发框架，可以以本项目为基础快速开发 Web API 服务，本项目大量参考 [singo](https://github.com/Gourouting/singo)，感谢！

## 项目工作

1. 整合了一些常用组件：

- [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的
- [GORM](http://gorm.io/docs/index.html): ORM工具，本项目需要配合Mysql使用
- [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端，用于缓存相关功能
- [godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具，方便使用环境变量
- [Jwt-Go](https://github.com/dgrijalva/jwt-go): Golang JWT组件，本项目使用基于 jwt 实现的 token 来做身份验证

2. 做了一个初步的模块划分：

- api文件夹就是MVC框架的controller，负责处理请求
- cache负责redis缓存相关的代码
- conf负责整个项目的静态配置
- middleware存放项目用到的相关中间件代码
- model文件夹负责存储数据库模型和数据库操作相关的代码
- router负责加载路由配置
- serializer负责响应码，响应格式等相关配置
- service负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
- util一些通用的小工具

3. 实现了一些常用代码方便参考和复用：

- 一个简单的用户模型
- /api/v1/user/register 用户注册接口
- /api/v1/user/login 用户登录接口
- /api/v1/user/me 用户资料接口(需传递token验证身份)
- /api/v1/user/logout 用户登出接口(需传递token验证身份)
- 一些可能用到的 util 小工具，目前有邮件发送，全局唯一ID生成器


## 使用教程

1. 下载项目到任意目录（除 GOPATH 路径中的 src 目录下，因为 Go Modules）
2. 修改项目文件夹名为你需要的项目名称
3. 进入项目目录，在终端执行 `go mod init 你需要的项目包名` 来修改项目包名
4. 修改项目所有文件中的 **go-api** 为 **第三步修改的包名** （可以用 goland 打开后 ctrl+shift+f 全局替换）
5. 项目依赖 MySQL 和 Redis，确保本机已经运行了这两个服务
6. 将 **example.env** 文件复制一份，重命名为 **.env** ，修改其中 MySQL 和 Redis 相关的配置
7. 执行 `go run main.go` 即可把项目跑起来啦
