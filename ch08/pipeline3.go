// goroutine管道.
// 限制channel的发送或接收.
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
	go counter(natu)
	// square
	go square(natu, squa)
	// printer(主goroutine)
	printer(squa)
}

// counter
// chan<- 中间不能有空格，限制channel只能用于发送
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

// square
// <-char，限制只有用于接收.
func square(in <-chan int, out chan<- int) {
	// 使用range迭代channel
	for x := range in {
		out <- x * x
	}
	close(out)
}

// printer
func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
