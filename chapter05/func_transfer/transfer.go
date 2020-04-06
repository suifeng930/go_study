package main

//函数参数的传递方式

// 值传递
// 引用传递

//  不管是值传递还是引用传递，传递给函数的都是变量的副本，不同的是，值传递的是值的拷贝，而引用传递的是地址的拷贝，一般来说
// 地址拷贝效率高，因为数据量小，而值拷贝决定于 拷贝的数据量的大小，数据越大，效率越低

// 值类型：  基本数据类型 int . float   bool  string  数组  结构体
//引用类型：  指针 slice   map   channel   interface 等都是引用类型

//值传递与引用传递使用特点

//      值类型默认是值传递， 变量直接存储值，内存通常在栈中分配
//      引用类型默认是引用传递，变量存储的是一个地址，这个地址指向的空间才是真正存储数据，内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对饮的数据空间就变成一个垃圾数据，会交由GC回收
func main() {

}
