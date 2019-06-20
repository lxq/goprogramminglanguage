// goroutine.
// 8.1 节
// @date 2019/6/20
// @author Allen Lin
// @email xqlin@qq.com

package main

import (
	"flag"
	"fmt"
	"time"
)

var n = flag.Int("n", 0, "输入要计算的第几个斐波那契数")

func main() {
	// 参数解析，必须!
	flag.Parse()

	// 输出表示主程序还在执行.
	go spinner(100 * time.Microsecond)

	fibN := fib(*n)
	fmt.Printf("\r第%d个斐波那契数是：%d\n", *n, fibN)
}

// 定时输出
func spinner(delay time.Duration) {
	for {
		for _, c := range `-\|/` {
			fmt.Printf("\r%c", c)
			time.Sleep(delay)
		}
	}
}

// 计算第n个斐波那契数
func fib(n int) int {
	if 2 > n {
		return n
	}
	return fib(n-1) + fib(n-2)
}
