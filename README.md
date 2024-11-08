
# go-blog

## 架构图

这也是一个MVC架构，api目录类似于controller层，通过router分发不同路由，然后在controller层这里调用service层，service层再调用dao层，dao层中进行数据库的操作。对于渲染页面的操作， views对应了MVC中的V，用来渲染页面，models是数据模型，定义了字段类型，以及和数据库字段的映射关系

如果你使用过nodejs框架nest.js，会发现这两者有略微的差别，nestjs中没有显式的用router来进行路由分发，而是在controller层中根据装饰器定义路由，如

```ts
  @Get('users')
  async findAll(): Promise<any> {
    return {
      code: 0,
      message: 'success',
      data: await this.userService.findAll()
    }
  }
```

一般来说， MVC中各角色的分工应该是

- c: controller层，也有的地方叫API层，作用：
  - 处理客户端的 HTTP 请求。
  - 将请求路由到相应的服务。
  - 对请求数据进行验证和处理。
  - 构建和返回 HTTP 响应。
  
  而service是在复杂业务场景下对于业务逻辑封装的抽象层，目的是为了保持controller逻辑更简洁，抽象出来的service也可以被多个controller调用,
    service的具体作用：
  - 处理业务逻辑。
  - 调用数据访问层（DAL）进行数据库操作。
  - 执行复杂的业务规则和数据处理。
  - 确保业务逻辑与 API 层解耦。

这里的结构更类似于express，或者eggjs项目的结构

```
+------------------+
|      api         |
| +--------------+ |
| |    common    | |
| | +----------+ | |
| | | config   | | |
| | | dao      | | |
| | | go.sum   | | |
| | | models   | | |
| | | router   | | |
| | | template | | |
| | | views    | | |
| | +----------+ | |
| +--------------+ |
+------------------+

+------------------+
|     config       | <-- 存储配置文件，例如数据库连接信息
+------------------+

+------------------+
|       dao        | <-- 数据访问层，包含与数据库交互的代码
+------------------+

+------------------+
|       models     | <-- 数据模型定义
+------------------+

+------------------+
|       router     | <-- 路由层，定义HTTP路由和处理函数
+------------------+

+------------------+
|      service     | <-- 服务层，包含业务逻辑
+------------------+

+------------------+
|      utils       | <-- 工具函数和通用实用程序
+------------------+

+------------------+
|      public      | <-- 静态资源（例如图片、CSS、JavaScript）
| +-------------+  |
| | resource    |  |
+------------------+

+------------------+
|     template     | <-- 存储HTML模板文件
+------------------+

+------------------+
|       views      | <-- 视图层，用于生成HTTP响应
+------------------+

+------------------+
|      main.go     | <-- 项目的入口点
+------------------+

```

## 运行

```bash
export DB_USER=你的mysql用户名
export DB_PASSWORD=你的mysql密码
export DB_HOST=你的mysql服务器地址
export DB_PORT=你的mysql端口
export DB_NAME=你的mysql数据库名

go run main.go
```

![image](https://github.com/user-attachments/assets/a752fafd-ca45-49af-a2d5-b02b25fe06e2)

然后打开http://localhost:5174

![image](https://github.com/user-attachments/assets/71922126-4334-4bf8-ad62-8813b0de3db9)

