# 抖音项目服务端 -- 第五届字节跳动青训营项目

### 一. 相关文档

- [抖音项目方案说明](https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof)
- [接口说明文档](https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707523)
- [极简抖音App使用说明](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)
- [服务端Demo仓库地址](https://github.com/RaymondCode/simple-demo)

### 二. 开发环境配置

1. 代码运行环境（版本号为本机环境，不需要完全相同）：

    - Golang 1.20.1

    - Mysql 8.0

    - Redis 5.0.14

2. 在`./config`文件夹下建立`config.yaml`文件，将以下代码复制到该文件中：

   ```yaml
   mysql:
     user: root
     pass: 123456
     host: 127.0.0.1
     port: 3306
     dbname: douyin
     charset: utf8mb4
     parsetime: True
     loc: Local
   
   redis:
     host : 127.0.0.1
     port : 6379
     pass :
   
   server:
     ip : 
     port : 8080
     msgport : 9090
   
   ffmpeg : C:\xxx\Douyin\middleware\ffmpeg\ffmpeg.exe
   ```

   根据自己的环境修改配置文件中的内容，一般只需要修改数据库的用户名和密码，服务器ip地址，ffmpeg.exe的绝对路径，若有其他需要也可自行修改。

3. 在mysql中建立名为douyin的数据库（可以是其他名字，需要修改`config.yaml`中的数据库名）。

4. 在终端输入`go run main.go msgServer.go`即可自动下载依赖并运行。

5. 使用安卓模拟器或安卓手机进行测试，[可以参考这篇文章](https://juejin.cn/post/7192600701745233979)。

### 三. 项目架构

1. 使用到的技术

    * 框架：gin、gorm

    * 数据库：Mysql

    * 其他
        * Redis：缓存
        * jwt：生成token、鉴权
        * bcrypt：对输入的password进行加密，数据库中存储加密后的密码
        * ffmpeg：截取视频的一帧作为封面
        * uuid：为上传的视频生成唯一的文件名
        * yaml：写配置文件

2. 采用 **repository → service → controller** 的分层结构：

    * **controller层**
        * 解析得到参数，传递给service层。

        * 如果需要返回数据信息，则调用service层的逻辑得到数据；如果不需要返回数据信息，只需要执行特定动作修改数据库，那么调用service层的逻辑执行这个动作。

        * 将得到的数据（如果有）与状态码和状态描述打包，返回响应。
    * **service层**

        * 如果上层需要返回数据信息，则进行参数检查、数据准备、数据打包；如果上层不需要返回数据信息，则进行参数检查、动作的执行。

        * 进行数据准备或动作执行时，需要调用repository层的逻辑。
    * **repository层**

        * 面向数据库进行增删改查。

3. 文件目录说明：其中controller和service文件夹中根据功能模块做了分包。

   ```
   Douyin 
   ├── /config/ 配置文件
   ├── /common/ 通用结构体
   ├── /controller/ 视图层
   ├── /service/ 逻辑层
   ├── /repository/ 数据层
   ├── /middleware/ 中间件
   │   ├── jwt/ 鉴权
   │   ├── ffmpeg/ 截取视频第一帧
   │   └── redis/ 缓存
   ├── /router/ 路由配置
   ├── /data/ 上传的视频文件存储在本地的路径，若不存在会自动创建
   ├── /go.mod/
   ├── msgServer.go  青训营demo中已经实现好的消息服务
   ├── main.go  程序入口
   └── README.md
   ```
