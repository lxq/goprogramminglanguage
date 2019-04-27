/**
 * hello world for Go.
 * 
 * @date 2019/4/26
 * @author Allen Lin
 * @email xqlin@qq.com
 */
package main


import(
	"fmt"
	"os"
)

 func main() {
	fmt.Println("你好， Go.")

	echo1()
	echo2()
 }

 /** 
  * test1 for os.Args .
 */
 func echo1() {
	const sep = " "
	var str string
	for i := 0; i < len(os.Args); i++ {
		str += sep + os.Args[i]
	}
	fmt.Println(str)
 }

 // 采用 for...rage，效率更高。
 func echo2() {
	// 短变更声明
	str, sep := "", ""
	// _ 为空标识符
	// for...range 一般用于字符串或Slice上进行迭代.
	// go 中只有一个循环：for [init;] [condition;] [post] {  }
	for _, arg := range(os.Args) {
		str += sep + arg
		sep = " "
	}
	fmt.Println(str)
 }