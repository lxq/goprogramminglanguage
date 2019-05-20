/**
 * dup练习.
 *
 * @date 2019/5/18
 * @author Allen Lin
 * @email xqlin@qq.com
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("UNIX uniq 指令.")
	// dup1()
	res := make(map[string]int)
	dupCounter(os.Stdin, res)

	for str, n := range res {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, str)
		}
	}
}

func dupCounter(f *os.File, res map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		res[input.Text()]++
	}
	// TODO: 忽略input.Err()中可能的错误.
}

// // 统计标准输入中次数大于1的行.
// func dup1() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}

// 	// TODO: 忽略input.Err()中可能的错误.
// 	for str, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, str)
// 		}
// 	}
// }
