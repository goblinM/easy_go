# Gin 使用
#### 基础命令以及使用
```
1.go env 查看go的环境配置
2.设置代理GOPROXY
    go env -w GOPROXY=https://goproxy.cn
3.go mod: 
    参考https://www.liwenzhou.com/posts/Go/go_dependency/
    go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
    go mod edit        编辑go.mod文件
    go mod graph       打印模块依赖图
    go mod init        初始化当前文件夹, 创建go.mod文件
    go mod tidy        增加缺少的module，删除无用的module
    go mod vendor      将依赖复制到vendor下
    go mod verify      校验依赖
    go mod why         解释为什么需要依赖
4.GOPATH 最好不要使用全局的，最好跟随项目
```

#### 路由：router
```
    原理： 将所有路由规则构造一颗前缀树
    1.基于httprouter开发的: 参考https://github.com/julienschmidt/httprouter
    2.支持Restful风格的API
    3.获取API的参数：
        Context的Param() 方法获取API参数
         router.GET("/user/:name/*action", func(c *gin.Context) {
            name := c.Param("name")
            action := c.Param("action")
            //截取/
            action = strings.Trim(action, "/")
            // 请求/http://localhost:8080/user/go/go.html
            // 输出：go is go.html
            c.String(http.StatusOK, name+" is "+action) // 
         })
    4.URL参数：
        通过DefaultQuery() 或者 Query() 方法获取
        eg: url为 user?name=zg
        router.get("/user", func(c *gin.Context){
            //http://localhost:8080/user 才会打印出来:默认值
            //http://localhost:8080/user?name=zg 才会打印出来:zg
            name = c.DefaultQuery("name", "默认值")
            fmt.Println(name)
        })
    5.表单参数：
        表单传输为post请求，http常见的传输格式为四种：
            application/json
            application/x-www-form-urlencoded
            application/xml
            multipart/form-data
        表单参数可以通过PostForm()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数
        router.POST("/form", func(c *gin.Context) {
            types := c.DefaultPostForm("type", "post")
            username := c.PostForm("username")
            password := c.PostForm("userpassword")
            // c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
            c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
        })
    6.上传文件
        单个文件：使用FormFile() 
            eg:
                file, err := context.FormFile("file") 或者
                // header里面可以获取到文件大小，类型
                _, headers, err := context.Request.FormFile("file") 
        多个文件：使用MultipartForm()
            eg:
                form, err := context.MultipartForm()
                files := form.File("files")
                for _, file := range files {}
        保存并上传文件使用：SaveUploadedFile()
            eg:
                context.SaveUploadedFile(file, file.Filename)
    7.router group: 管理一些相同的URL
        r := gin.Default()
        v1 := r.Group("/v1")
        {
            v1.GET("/login", login)
            v1.GET("/submit", submit)
        }
        v2 := r.Group("/v2")
        {
            v2.POST("/login", login)
            v2.POST("/submit", submit)
        }
        r.Run(":8000")
        func login(c *gin.Context) {
           name := c.DefaultQuery("name", "jack")
           c.String(200, fmt.Sprintf("hello %s\n", name))
        }
        
        func submit(c *gin.Context) {
           name := c.DefaultQuery("name", "lily")
           c.String(200, fmt.Sprintf("hello %s\n", name))
        }
        // curl http://localhost:8000/v1/submit -X POST
        // curl http://localhost:8000/v2/submit -X POST
    8.路由拆分以及注册
        func ResponseHandle(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "message": "ok"
            }
        }
        8.1 基本的路由注册（适合demo）
            // main.go 中使用
            r := gin.Default()
            r.GET("/", ResponseHandle)
        8.2 路由拆分成单独文件或包
            // 新建 routers.go 文件，在这里面注册路由
            func SetupRouter() *gin.Engine {
                r := gin.Default()
                r.GET("/hello", ResponseHandle)
                return r
            }
            // 在main.go 文件中调用 routers.go 中的 SetupRouter()
            r := setupRouter()
        8.3 路由拆分成多个文件
            // blog.go 将blog相关的路由注册到指定的路由器
            func LoadBlog(e *gin.Engine) {
                e.GET("/blog", ResponseHandle)
            }
            // shop.go 将shop相关的路由注册到指定的路由器
            func LoadShop(e *gin.Engine) {
                e.GET("/shop", ResponseHandle)
            }
            // main.go
            r := gin.Default()
            router.LoadBlog(r)
            router.LoadShop(r)
        8.4 路由拆分到不同的APP
            在项目目录下单独定义一个app目录，用来存放不同业务线的代码文件，这样就很容易进行横向扩展。
            类似：
            gin_demo
            ├── app
            │   ├── blog
            │   │   ├── handler.go
            │   │   └── router.go
            │   └── shop
            │       ├── handler.go
            │       └── router.go
            ├── go.mod
            ├── go.sum
            ├── main.go
            └── routers
                └── routers.go
            
            app/blog/router.go用来定义blog相关路由信息
            app/blog/handler.go用来定义blog相关处理逻辑
            app/shop/router.go用来定义shop相关路由信息
            app/shop/handler.go用来定义shop相关处理逻辑
            
            routers/routers.go中根据需要定义Include函数用来注册子app中定义的路由
            Init函数用来进行路由的初始化操作
                type Option func(*gin.Engine)
                var options = []Option{}
                // 注册app的路由配置
                func Include(opts ...Option) {
                    options = append(options, opts...)
                }
                // 初始化
                func Init() *gin.Engine {
                    r := gin.New()
                    for _, opt := range options {
                        opt(r)
                    }
                    return r
                }
            main.go中按如下方式先注册子app中的路由，然后再进行路由的初始化
                // 加载多个APP的路由配置
                routers.Include(shop.Routers, blog.Routers)
                // 初始化路由
                r := routers.Init()
                if err := r.Run(); err != nil {
                    fmt.Println("startup service failed, err:%v\n", err)
                }    
```     
#### 数据解析和绑定
```
    1.JSON数据解析与绑定
        绑定：context.ShouldBindJSON(&json)
        // gin.H封装了生成json数据的工具
        context.JSON(http.StatusOK, gin.H{"status": "200"})
    2.表单数据解析与绑定
        绑定：context.Bind(&form)
        context.JSON(http.StatusOK, gin.H{"status": "200"})
    3.URI数据解析与绑定
        绑定：context.ShouldBindUri(&login)
        context.JSON(http.StatusOK, gin.H{"status": "200"})
```
#### 数据渲染
```
    1.数据格式响应
        1.1 json格式：gin.H
             c.JSON(200, gin.H{"message": "someJSON", "status": 200})
        1.2 struct 结构体格式
            var msg struct {
                 Name string
            }
            msg.Name = "hello"
            c.JSON(200, msg)
        1.3 xml格式
            c.XML(200, gin.H{"message": "abc"})
        1.4 yaml格式
            c.YAML(200, gin.H{"name": "zhangsan"})
        1.5 protobuf格式,谷歌开发的高效存储读取的工具
            r.GET("/someProtoBuf", func(c *gin.Context) {
                reps := []int64{int64(1), int64(2)}
                // 定义数据
                label := "label"
                // 传protobuf格式数据
                data := &protoexample.Test{
                    Label: &label,
                    Reps:  reps,
                }
                c.ProtoBuf(200, data)
            })
    2.html模板渲染：
        参考：https://www.liwenzhou.com/posts/Go/go_template/
            http://www.topgoer.com/gin框架/gin渲染/html模板渲染.html
        gin支持加载HTML模板, 然后根据模板参数进行配置并返回相应的数据，本质上就是字符串替换
        LoadHTMLGlob()方法可以加载模板文件
        eg:
        r := gin.Default()
        // tem 文件下是存放html文件
        r.LoadHTMLGlob("tem/*")
        r.GET("/index", func(c *gin.Context) {
            c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
        })
        r.Run()
    3.重定向：
        c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
    4.同步异步：
         goroutine机制可以方便地实现异步处理
        另外，在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本
        异步是在context里面增加一个类似go func() {}内置函数，把处理的逻辑放在里面
            r := gin.Default()
            // 1.异步
            r.GET("/long_async", func(c *gin.Context) {
                // 需要搞一个副本
                copyContext := c.Copy()
                // 异步处理
                go func() {
                    time.Sleep(3 * time.Second)
                    log.Println("异步执行：" + copyContext.Request.URL.Path)
                }()
            })
            // 2.同步
            r.GET("/long_sync", func(c *gin.Context) {
                time.Sleep(3 * time.Second)
                log.Println("同步执行：" + c.Request.URL.Path)
            })
        
            r.Run(":8000")
```
