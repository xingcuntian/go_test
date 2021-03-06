package main

import (
	"fmt"
	"sync/atomic"
)
//函数提供的原子操作共有五种：增或减，比较并交换，载入，存储和交换
// 增或减
func main(){
	var i32 int32
	fmt.Println("=====old i32 value=====")
    fmt.Println(i32)
	  //第一个参数值必须是一个指针类型的值,因为该函数需要获得被操作值在内存中的存放位置,以便施加特殊的CPU指令
    //结束时会返回原子操作后的新值
	newI32 := atomic.AddInt32(&i32,3)
    fmt.Println("=====new i32 value=====")
    fmt.Println(i32)
    fmt.Println(newI32)

	var i64 int64
	fmt.Println("================old i64 value=====")
	fmt.Println(i64)
	newI64 := atomic.AddInt64(&i64,-3)
	fmt.Println("===========new i64 value ==========")
	fmt.Println(i64)
	fmt.Println(newI64)
}
// =====old i32 value=====
// 0
// =====new i32 value=====
// 3
// 3
// ================old i64 value=====
// 0
// ===========new i64 value ==========
// -3
// -3