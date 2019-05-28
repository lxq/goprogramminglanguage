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

	// slice的定义和初始化，区别于数组是：没有指定长度.
	s := []int{111, 222, 333, 444, 555}
	reverse(s)
	fmt.Println(s)

	testAppend()

	x := make([]int, 2, 2)
	fmt.Printf("len(x)=%d, cap(x)=%d\n", len(x), cap(x))
	x = appendInt(x, 123)
	fmt.Printf("len(x)=%d, cap(x)=%d, slice:%v\n", len(x), cap(x), x)

}

// reverse a slice
func reverse(s []int) {
	// 这里不能用i++,j--，原因是i++,j--是语句，不是表达式
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		// for 和下面语句都体现了go的多赋值便利性
		s[i], s[j] = s[j], s[i]
	}
}

func testAppend() {
	var runes []rune
	for _, c := range "Hello, 你好!" {
		runes = append(runes, c)
	}
	fmt.Printf("%q\n", runes)
}

// 往int slice尾部append一个int，以理解slice原理。
func appendInt(s []int, y int) []int {
	var ss []int
	l := len(s) + 1
	if l <= cap(s) {
		// 即原slice容量还可以容纳至少一个int
		ss = s[:l]
	} else {
		// 原slice容量不够，按翻倍进行扩展
		c := l
		if c < 2*len(s) {
			// TODO:只有s []int为空时，才会出现这种情况.
			c = 2 * len(s)
		}
		// 分配新的slice空间
		ss = make([]int, l, c)
		// copy原slice数据
		copy(ss, s)
	}
	ss[l-1] = y
	return ss
}
