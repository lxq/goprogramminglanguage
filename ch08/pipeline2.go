// goroutine管道.
// 通过range迭代channel.
// 8.4.2 节
// @date 2019/6/20
// @author Allen Lin
// @email xqlin@qq.com

package main

import "fmt"

func main() {
	natu := make(chan int)
	squa := make(chan int)

	// counter
	go func() {
		for x := 0; x < 100; x++ {
			natu <- x
		}
		close(natu)
	}()

	// square
	go func() {
		// 使用range迭代channel
		for x := range natu {
			squa <- x * x
		}
		close(squa)
	}()

	// printer(主goroutine)
	for x := range squa {
		fmt.Println(x)
	}
}
