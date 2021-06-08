# 基础go入门

#### 基础组成结构
```
    1.包声明（package main）  
    2.引入包（import "fmt"）
    3.函数定义（func main(){}）
    4.变量定义以及应用
    5.语句 以及表达式
        a :=1  
        b:= 10  
        a,b=b,a
    6.注释（多行注释：/**注释*/   单行注释：// 注释）
    
```

#### 基础数据类型快速了解
```
    1.布尔类型：true,false
    2.数字类型：整型，浮点型
    3.字符串类型： 使用utf-8编码的标识的Unicode文本
    4.派生类型：
        4.1 指针类型： & 取指针地址   * 取指针的值
        4.2 数组类型：
        4.3 结构化类型：struct
        4.4 Channel 类型
        4.5 函数类型： func test(){}
        4.6 切片类型：
        4.7 接口类型（interface）
        4.8 Map类型
```


#### 基础语法快速了解
```
    1.go 标记： 可以是关键字，标识符，常量，字符串，符号等组成
    2.行分隔符： 一行代表一个语句结束；多个语句合并一行，使用;人为区分（开发不建议使用）
    3.注释（上面已列出）
    4.标识符：一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列
    5.字符串连接：字符串可以通过 + 实现
    6.关键字：
        常用的：
            var func const interface package import case for 
            if else return go goto continue case default swith
            struct type range map chan fallthrough defer
        还有一些没列出来的可以自行百度
    7.fmt.Sprintf(): 格式化字符串并赋值给新串
         func TestFormatString() {
        	// Sprintf 格式化字符串并赋值给新串
        	name := "Mary"
        	age := 20
        	word := "my name is %s, i'm %d years old!"
        	var introduce = fmt.Sprintf(word, name, age)
        	fmt.Println(introduce)
         }
    8.变量：字母，数字，下划线组成，首个字母不能为数字
        8.1 var 声明变量,可一次声明多个变量：
            var name1,name2 string = "mary", "jane"
        8.2 变量声明   
            第一种，指定变量类型，如果没有初始化，则变量默认为零值。 
            不同类型的零值：
                数值类型为 0
                布尔类型为 false
                字符串为  ""
                下面这些是为 nil:
                    var a *int
                    var a []int
                    var a map[string] int
                    var a chan int
                    var a func(string) int
                    var a error // error 是接口
            第二种，根据值自行判定变量类型。
        8.3 值类型和引用类型
            值类型，使用这些类型的变量直接指向存在内存中的值
            当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝
    9.常量：不会被修改的量
        const
        iota:  
            特殊常量，可以认为是一个可以被编译器修改的常量。
            iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
            iota 可以被用作枚举值：
                const (
                        a = iota   //0
                        b          //1
                        c          //2
                        d = "ha"   //独立值，iota += 1
                        e          //"ha"   iota += 1
                        f = 100    //iota +=1
                        g          //100  iota +=1
                        h = iota   //7,恢复计数
                        i          //8
                    )
    10.语句
        条件语句：
            if 表达式 {} else {}
            switch 语句：switch 语句用于基于不同条件执行不同动作。
                switch var1 {
                    case val1 :
                        ...
                    case val2:
                        //使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
                        ...
                        fallthrough
                    default:
                        ...
                }
            select 语句：类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
                select {
                    case communication clause  :
                       statement(s);      
                    case communication clause  :
                       statement(s);
                    /* 你可以定义任意数量的 case */
                    default : /* 可选 */
                       statement(s);
                }
        循环语句：
            for 循环：
                for init; condition; post {}
                for condition {}
                for {}
                
                init： 一般为赋值表达式，给控制变量赋初值；
                condition： 关系表达式或逻辑表达式，循环控制条件；
                post： 一般为赋值表达式，给控制变量增量或减量。
    11.语言函数
        func function_name( [parameter list] ) [return_types] {
           函数体
        }
        值传递：值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
        引用传递：引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
        内置函数：函数定义后可作为另外一个函数的实参数传入
        闭包：闭包是匿名函数，可在动态编程中使用
        方法：方法就是一个包含了接受者的函数
        
```

#### 基础数据类型快速使用
```
    1.数组：
        语法格式：var array_name [size] value_type
                 size 不确定的时候可以使用 ...  代替
                 eg: var fruit_array [5] string
        初始化数组： 
                var fruit_array = [5] string {'apple', 'orange', 'watermelon', 'pear', 'lemon'}
                animal_array = [...] string {'cat', 'dog', 'bird', 'fox', 'panda', 'snake'}
                初始化中的{} 元素个数 不能大于 [] 中的数字
    
    2.指针：
        & 取指针地址   * 取指针的值
        声明格式： var var_name *var-type
                  var ip *int // 指向整形
                  var name *string  //指向字符
        指针使用流程：
            2.1 定义指针变量
            2.2 为指针变量赋值
            2.3 访问指针变量中指向地址的值
            eg: 
                a:= 20
                var ip *int  // 声明指针变量
                ip = &a      // 指针变量的存储地址
                fmt.Printf("a 变量的地址是: %x\n", &a)
                fmt.Printf("ip 变量存储的指针地址是: %x\n", ip)
                /* 使用指针访问值 */
                fmt.Printf("*ip 变量的值: %d\n", *ip )
        空指针：当一个指针被定义后没有分配大任何变量，它的值是nil(空指针)
            指针变量通常缩写成 ptr
            空指针判断：if (ptr != nil){}
        指针数组：
        指向指针的指针：一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
            声明格式： var prt **int
            var a int
            var ptr *int
            var pptr **int
            a = 3000
            /* 指针 ptr 地址 */
            ptr = &a
            /* 指向指针 ptr 地址 */
            pptr = &ptr
            /* 获取 pptr 的值 */
            fmt.Printf("变量 a = %d\n", a )
            fmt.Printf("指针变量 *ptr = %d\n", *ptr )
            fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

    3.结构体
        结构体格式：
            type struct_name struct {
                member definition
                ...
                member definition
            }
        结构体用于变量声明，语法格式如下：
            variable_name := struct_name {value1, value2...valuen}
        结构体指针：
            var struct_pointer *Books
            定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前：
            struct_pointer = &Book1
            使用结构体指针访问结构体成员，使用 "." 操作符：
            struct_pointer.title
    4.语言切片slice：
        语言切片是对数组的抽象
        灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
        定义切片：
            var slices_name []type
        使用 make() 函数来创建切片:
            slices_name := make([]type, len)
            len 是数组的长度并且也是切片的初始长度。
        也可以指定容量，其中 capacity 为可选参数。
            make([]type, len, capacity)
        切片初始化：
            s := [] int {1,2,3}
            [] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。
        len(): 获得长度
        cap(): 测量切片最长可以达到多少
        空切片：一个切片在未初始化之前默认为 nil，长度为 0
        切片截取:通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]
        append(): 添加元素  append(slice_name, 11)
        copy(): 拷贝元素 copy(slice_name1, slice_name)
        
    5.语言范围：range
        range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
        在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
    6.语言集合：map
        无序的键值对的集合
        通过key来快速检索数据
        定义map:
            使用内建函数make 或者使用map 关键字来定义
            如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
            /*声明变量，默认map 是 nil*/
            var map_variable map[key_data_type]value_data_type
            /*使用make 函数*/
            map_variable := make(map[key_data_type]value_data_type)
        
            eg:
                var countryCapitalMap map[string]string  /*创建集合*/
                countryCapitalMap = make(map[string]string)
                capitalNameArray := []string{"巴黎", "罗马", "东京", "新德里"}
                countryNameArray := []string{"France", "Italy", "Japan", "India"}
                for index, country := range countryNameArray {
                    countryCapitalMap[country] = capitalNameArray[index]
                }
                for country := range countryCapitalMap {
                    fmt.Println(country, "首都是", countryCapitalMap[country])
                }

        delete() 函数：
            delete()函数用于删除集合的元素，参数为map和其对应的key
            delete(countryCapitalMap, "France")
    7.语言接口：interface
        把所有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口
        /*定义接口*/
        type Animal interface {
        	// 接口方法
        	colorType()
        }
        /*定义结构体*/
        type Cat struct {
        	color string
        }
        /**实现接口方法*/
        func (cat Cat) colorType() {
       	    fmt.Println("color = ", cat.color)
        }
        /*调用*/
        func main() {
        	var animal Animal
            animal := Cat{color:"white"}
            animal.colorType()
        }
    8.错误处理 
        通过内置的错误接口提供了非常简单的错误处理机制
        定义：
            type error interface {
                Error() string
            }   
    9.go并发
        支持并发，只需要通过go 关键字来开启goroutine 
        goroutine 是轻量级线程，goruntine的调度是由golang运行时进行管理的
        Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。 
        同一个程序中的所有 goroutine 共享同一个地址空间。
        goruntine 语法格式：
            go 函数名(参数列表)
            eg: go f(x,y,z)
        
        channel: 通道
            通道（channel）是用来传递数据的一个数据结构。
            通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。
            操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
            ch <- v    // 把 v 发送到通道 ch
            v := <-ch  // 从 ch 接收数据
                       // 并把值赋给 v
            声明通道：chan关键字
                通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：
                ch := make(chan int, 100)
        Go 遍历通道与关闭通道
        Go 通过 range 关键字来实现遍历读取到的数据，类似于与数组或切片
            v, ok := <-ch
        如果通道接收不到数据后 ok 就为 false，这时通道就可以使用 close() 函数来关闭。
            
```