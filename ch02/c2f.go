/**
 * 练习调用自定义的包.
 * 2.6 节
 *
 * @date 2019/5/23
 * @author Allen Lin
 * @email xqlin@qq.com
 */
package main

import (
	"fmt"
	"gopl/ch02/tmpc"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if nil != err {
			fmt.Fprintf(os.Stderr, "参数错误:%v\n", err)
			os.Exit(1)
		}

		f := tmpc.Fahrenheit(t)
		c := tmpc.Celsuis(t)
		// 下面"%s"会自动调用tmpc包中定义的String()方法进行转换.
		fmt.Printf("%s = %s, %s = %s\n", f, tmpc.FtoC(f), c, tmpc.CtoF(c))
	}
}
