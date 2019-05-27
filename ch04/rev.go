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
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reverse(a[:])
	fmt.Println(a)
}

// reverse a slice
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
