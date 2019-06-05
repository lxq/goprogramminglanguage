// Go 函数练习.
// 5.7 节
// @date 2019/6/5
// @author Allen Lin
// @email xqlin@qq.com

package main

import (
	"fmt"
)

func main() {
	fmt.Println("第五章 函数练习.")
	fmt.Printf("sum()=%d, sum(3)=%d, sun(1,2,3,4)=%d\n", sum(), sum(3), sum(1, 2, 3, 4))
}

// 可变参数
func sum(arr ...int) int {
	t := 0
	for _, v := range arr {
		t += v
	}
	return t
}
