# Go MySQL
#### MySQL快速使用
```
    使用第三方开源的mysql库: 
        github.com/go-sql-driver/mysql （mysql驱动） 
        github.com/jmoiron/sqlx （基于mysql驱动的封装）

    1.安装：
        go get -u github.com/go-sql-driver/mysql
        go get -u  github.com/jmoiron/sqlx 
    2.连接数据库：
        database, err := sqlx.Open("mysql", "root:XXXX@tcp(127.0.0.1:3306)/test")
        //database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
    3.mysql事务特性（ACID）：
        1.原子性
        2.一致性
        3.隔离性
        4.持久性
    4.MySQL事务应用：
        1） import (“github.com/jmoiron/sqlx")
        2)  Db.Begin()        开始事务
        3)  Db.Commit()        提交事务
        4)  Db.Rollback()     回滚事务
```
#### gorm 概览信息
```
    1.全特性 ORM (几乎包含所有特性)
    2.模型关联 (一对一， 一对多，一对多（反向）， 多对多， 多态关联)
    3.钩子 (Before/After Create/Save/Update/Delete/Find)
    4.预加载
    5.事务
    6.复合主键
    7.SQL 构造器
    8.自动迁移
    9.日志
    10.基于GORM回调编写可扩展插件
    11.全特性测试覆盖
    12.开发者友好
```

#### gorm 快速使用
```
    1.安装： go get -u github.com/jinzhu/gorm
    2.使用步骤：
        1.import gorm
        2.定义数据库结构(模型)：
            type Product struct {
                gorm.Model
                Code string
                ...
            }
        3.操作数据库基本操作
            db := gorm.Open()： 连接数据库
            defer db.Close() ： 关闭数据库连接
            db.AutoMigrate(&Product{}) ： 自动检查Product结构是否变化，变化则进行迁移
            db.Create(&Product{Code:"a"}) ： 增
            var product Product
            db.Frist(&product, 1) ： 查找id为1的产品
            db.Model(&product).Update("Code", "dd") ： 改
            db.Delete(&product) ： 删除
            // 为插入 SQL 语句添加额外选项
            db.Set("gorm:insert_option", "ON CONFLICT").Create(&product)

        4.惯例
            gorm.Model 是一个包含一些基本字段的结构体, 
            包含的字段有 ID，CreatedAt， UpdatedAt， DeletedAt
        5.连接数据库
            import (
              "github.com/jinzhu/gorm"
              _ "github.com/jinzhu/gorm/dialects/mysql"
            )
            
            func main() {
              db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
              defer db.Close()
            }
    3.CRUD接口
        1.创建
            user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
            db.Create(&user)
        2.查询: http://www.topgoer.com/数据库操作/gorm/CRUD接口/查询.html
            2.1 基础查询
                // 获取第一条记录，按主键排序
                db.First(&user)
                // 获取一条记录，不指定排序
                db.Take(&user)
                // 获取最后一条记录，按主键排序
                db.Last(&user)
                // 获取所有的记录
                db.Find(&users)
                // 通过主键进行查询 (仅适用于主键是数字类型)
                db.First(&user, 10)
            2.2 where 查询
                // 获取第一条匹配的记录
                db.Where("name = ?", "jinzhu").First(&user)
                // SELECT * FROM users WHERE name = 'jinzhu' limit 1;
                // 获取所有匹配的记录
                db.Where("name = ?", "jinzhu").Find(&users)
                // SELECT * FROM users WHERE name = 'jinzhu';
            2.3 Struct & Map
                当通过struct进行查询的时候，GORM 将会查询这些字段的非零值， 
                意味着你的字段包含 0， ''， false 或者其他 零值, 将不会出现在查询语句中
                你可以考虑适用指针类型或者 scanner/valuer 来避免这种情况。
                // Struct
                db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
                //// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;
                // Map
                db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
                //// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
                // 多主键 slice 查询
                db.Where([]int64{20, 21, 22}).Find(&users)
                //// SELECT * FROM users WHERE id IN (20, 21, 22);
```