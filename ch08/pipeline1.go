// goroutine管道.
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
		for x := 0; ; x++ {
			natu <- x
		}
	}()

	// square
	go func() {
		for {
			x := <-natu
			squa <- x * x
		}
	}()

	// printer(主goroutine)
	for {
		fmt.Println(<-squa)
	}
}
