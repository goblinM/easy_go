语言切片(Slice)
```
    语言切片slice：
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
    
    语言范围：range
        range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
        在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
        
```