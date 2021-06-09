# Gin 快速使用
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
#### 中间件
```
    1.全局中间件
        所有的请求都会经过此中间件，再去到对应的接口
        中间件写法：
        // middleware.go 写入
        func MiddleWare() gin.HandlerFunc {
            return func(c *gin.Context) {
                //业务处理逻辑
            }
        }
        // main.go
        // 注册中间件：没有进行路由拆分就在main.go, 如果拆分了后在routers.go中
        r := gin.Default()
        r.Use(MiddleWare())
    2.Next()方法
        // middleware.go 中写入
        c.Next()
    3.局部中间件
        定义好一个局部中间件函数
        然后直接在路由注册的时候直接使用，则可当成局部中间件使用
        //
        func PartMiddleWare() gin.HandlerFunc {
            return func(c *gin.Context) {
                //业务处理逻辑
            }
        }
        // main.go
        r := gin.Default()
        r.GET("/test", PartMiddleWare(), func(c *gin.Context){})
```
#### 会话控制
```
    1.cookie
        介绍：
            Cookie是解决HTTP协议无状态的方案之一
            Cookie实际上就是服务器保存在浏览器上的一段信息。浏览器有了Cookie之后，每次向服务器发送请求时都会同时将该信息发送给服务器，服务器收到请求后，就可以根据该信息处理请求
            Cookie由服务器创建，并发送给浏览器，最终由浏览器保存
        用途：测试服务端发送cookie给客户端，客户端请求时携带cookie
        获取： context.Cookie("key_cookie")
        设置： 
            // 给客户端设置cookie
            //  maxAge int, 单位为秒
            // path,cookie所在目录
            // domain string,域名
            // secure 是否智能通过https访问
            // httpOnly bool  是否允许别人通过js获取自己的cookie
            c.SetCookie("key_cookie", "value_cookie", 60, "/",
                "localhost", false, true)
            }
        缺点：
            不安全，明文
            增加带宽消耗
            可以被禁用
            cookie有上限
    2.session
        官网地址：http://www.gorillatoolkit.org/pkg/sessions
        简单示例：/easy_go/ginDemo/app/middlewares/session.go
        gorilla/sessions为自定义session后端提供cookie和文件系统session以及基础结构
        主要功能：
            1.可用于设置签名
            2.内置的后端可将session存储在cookie或文件系统中。
            3.Flash消息：一直持续读取的session值。
            4.切换session持久性（又称“记住我”）和设置其他属性的便捷方法。
            5.旋转身份验证和加密密钥的机制。
            6.每个请求有多个session，即使使用不同的后端也是如此。
            7.自定义session后端的接口和基础结构：可以使用通用API检索并批量保存来自不同商店的session。
```
#### 参数验证
``` 
    1.结构体验证
        type Person struct {
            //不能为空并且大于10
            Age      int       `form:"age" binding:"required,gt=10"`
        }
        func main() {
            r := gin.Default()
            r.GET("/struct_auth", func(c *gin.Context) {
                var person Person
                if err := c.ShouldBind(&person); err != nil {
                    c.String(500, fmt.Sprint(err))
                    return
                }
                c.String(200, fmt.Sprintf("%#v", person))
            })
            r.Run()
        }
    2.自定义验证:
        参考：http://www.topgoer.com/gin框架/参数验证/自定义验证.html
        // 1、自定义的校验方法
        func nameNotNullAndAdmin(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
            if value, ok := field.Interface().(string); ok {
                // 字段不能为空，并且不等于  admin
                return value != "" && !("5lmh" == value)
            }
            return true
        }
         // 3、将我们自定义的校验方法注册到 validator中
        if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
            // 这里的 key 和 fn 可以不一样最终在 struct 使用的是 key
            v.RegisterValidation("NotNullAndAdmin", nameNotNullAndAdmin)
        }
    3.多语言翻译验证
```
#### 日志文件
```
    gin.DisableConsoleColor()
    // Logging to a file.
    f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f)
    // 如果需要同时将日志写入文件和控制台，请使用以下代码。
    // gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
```
#### 实时加载
```
    参考：https://www.jianshu.com/p/d7916f21d38c
    1.Air:
        Air能够实时监听项目的代码文件，在代码发生变更之后自动重新编译并执行，大大提高gin框架项目的开发效率
        Air 特性：    
            1.彩色日志输出
            2.自定义构建或二进制命令
            3.支持忽略子目录
            4.启动后支持监听新目录
            5.更好的构建过程
        安装： go get -u github.com/cosmtrek/air
        
    2.Fresh
        github地址：https://github.com/gravityblast/fresh
        Fresh是一个命令行工具，每次保存Go或模版文件时，该工具都会生成或重新启动Web应用程序。
        Fresh将监视文件事件，并且每次创建/修改/删除文件时，Fresh都会生成并重新启动应用程序。
        如果go build返回错误，它会将记录在tmp文件夹中。
        安装：go get -u github.com/pilu/fresh
        使用：进入你的项目目录 然后执行 fresh
    3.bee
        github地址：https://github.com/beego/bee
        安装：go get -u github.com/beego/bee
        使用：进入你的项目目录 然后执行 bee run
    4.gowatch 
        github地址：https://github.com/silenceper/gowatch
        安装：go get github.com/silenceper/gowatch
        使用： gowatch -o ./bin/demo -p ./cmd/demo
            -o : 非必须，指定build的目标文件路径
            -p : 非必须，指定需要build的package（也可以是单个文件）
            -args: 非必须，指定程序运行时参数，例如：-args='-host=:8080,-name=demo'
            -v: 非必须，显示gowatch版本信息
    5.gin
        github地址:https://github.com/codegangsta/gin
        安装： go get github.com/codegangsta/gin
        使用：gin run main.go    
    6.realize
        github地址：https://github.com/oxequa/realize
        安装：go get github.com/oxequa/realize
             或者  GO111MODULE=off go get github.com/oxequa/realize
        使用：
            # 首先进行初始化 默认配置即可
            $ realize init
            # 执行项目
            $ realize start
            # 添加命令
            $ realize add
            # 删除命令
            $ realize init
```