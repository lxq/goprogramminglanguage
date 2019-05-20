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

	res := make(map[string]int)

	fileNames := os.Args[1:]
	if len(fileNames) == 0 {
		fmt.Println("未输入文件，采用命令行输入进行统计!")
		dupCounter(os.Stdin, res)
	} else {
		for _, arg := range fileNames {
			f, err := os.Open(arg)
			if nil != err {
				fmt.Fprintf(os.Stderr, "dup files: %v\n", err)
				continue
			}
			dupCounter(f, res)
			f.Close()
		}
	}

	for str, n := range res {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, str)
		}
	}
}

// 把File、os.Stdin进行统一.
func dupCounter(f *os.File, res map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		res[input.Text()]++
	}
	// TODO: 忽略input.Err()中可能的错误.
}
