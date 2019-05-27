/**
 * array& slice
 * 4.2 节
 *
 * @date 2019/5/28
 * @author Allen Lin
 * @email xqlin@qq.com
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array & Slice.")
	// a 是数组, [...]表示长度由初始化值个数确定.
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reverse(a[:])
	fmt.Println(a)
}

// reverse a slice
func reverse(s []int) {
	// 这里不能用i++,j--，原因是i++,j--是语句，不是表达式
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		// for 和下面语句都体现了go的多赋值便利性
		s[i], s[j] = s[j], s[i]
	}
}
