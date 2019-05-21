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
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("UNIX uniq 指令.")

	res := make(map[string]int)

	fileNames := os.Args[1:]
	if len(fileNames) == 0 {
		fmt.Println("未输入文件，采用命令行输入进行统计!")
		dupStream(os.Stdin, res)
	} else {
		for _, arg := range fileNames {
			f, err := os.Open(arg)
			if nil != err {
				fmt.Fprintf(os.Stderr, "dup files: %v\n", err)
				continue
			}
			dupStream(f, res)
			f.Close()
			// read a file into memroy
			dupFile(arg)
		}
	}

	fmt.Println("全部统计结果:")
	for str, n := range res {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, str)
		}
	}
}

// 把File、os.Stdin进行统一.
// 这种逐行处理相当于“流”模式.
func dupStream(f *os.File, res map[string]int) {
	input := bufio.NewScanner(f)
	// Scanner.Scan()读取下一行并去除行尾换行符。
	for input.Scan() {
		res[input.Text()]++
	}
	// TODO: 忽略input.Err()中可能的错误.
}

/**
 * 统计单个文件中的重复行次数.
 * 用于练习ioutil.ReadFile(), string.Split()等方法。
 */
func dupFile(filename string) {
	fmt.Println("统计单个文件中的行重复次数.")
	// TODO: 加入文件是否存在的判断.
	res := make(map[string]int)
	// ReadFile()一次把全部内容读取内存.
	// 返回值 data 是一个可以转换为字符串的字节slice.
	// TODO: 返回值为字节slice，是否意味着ReadFile()不能读取二进制文件？
	data, err := ioutil.ReadFile(filename)
	if nil != err {
		fmt.Fprintf(os.Stderr, "%s 读取错误: %v\n", filename, err)
		return
	}

	// strings.Split()按给定分隔符把字符串分成slice.
	for _, line := range strings.Split(string(data), "\n") {
		// TODO: 这里读取的行内容，并未考虑换行符情况，最行一行内容没有换行则会出现即使其他行内容一致也会统计成不一样.
		res[line]++
	}
	for line, n := range res {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
