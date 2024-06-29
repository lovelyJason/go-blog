# go-blog

## 架构图
这也是一个MVC架构，api目录类似于controller层，通过router分发不同路由，然后在controller层这里调用service层，service层再调用dao层，dao层中进行数据库的操作。对于渲染页面的操作， views对应了MVC中的V，用来渲染页面，models是数据模型，定义了字段类型，以及和数据库字段的映射关系

如果你使用过nodejs框架nest.js，会发现这两者有略微的差别，nestjs中没有显式的router区进行路由分发，而是在controller层中根据装饰器定义路由，如

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

这里的结构更类似于express，或者eggjs项目的结构， 只是编程语言的不同，设计理念都是差不多

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
