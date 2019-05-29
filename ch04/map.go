/**
 * map example
 * 4.3 节
 *
 * @date 2019/5/29
 * @author Allen Lin
 * @email xqlin@qq.com
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

func main() {
	fmt.Println("map exaples.")

	// map键值必须能通过 == 进行比较.
	ages := make(map[string]int)
	ages["lxq"] = 40
	ages["lmz"] = 38
	ages["lry"] = 9
	fmt.Printf("%q\n", ages)
	// map元素的迭代顺序不固定
	for k, v := range ages {
		fmt.Printf("k=%s : v=%d\t", k, v)
	}
	fmt.Println()
	// sort.Strings()对字符串键值map进行排序.
	names := make([]string, 0, len(ages))
	// 下面for直接忽略值.
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("k=%s : v=%d\t", name, ages[name])
	}

	fmt.Println()
	countUnicode()
}

func countUnicode() {
	// 字符数量
	mapNum := make(map[rune]int)
	// utf-8编码长度
	var utflen [utf8.UTFMax + 1]int
	// 无效字符数量
	invNum := 0

	inReader := bufio.NewReader(os.Stdin)
	for {
		// rune, nbytes, error
		r, n, err := inReader.ReadRune()
		if io.EOF == err {
			break
		}
		if nil != err {
			fmt.Fprintf(os.Stderr, "读取错误：%v\n", err)
			os.Exit(1)
		}
		if 1 == n && unicode.ReplacementChar == r {
			invNum++
			continue
		}
		mapNum[r]++
		utflen[n]++
	}

	fmt.Println("rune\t数量\n")
	for c, n := range mapNum {
		// %q: 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
		// %v:值的默认格式表示。当输出结构体时，扩展标志（%+v）会添加字段名
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Println("长度\t数量\n")
	for i, n := range utflen {
		fmt.Printf("%d\t%d\n", i, n)
	}
	fmt.Printf("无效utf8字符数量:%d\n", invNum)
}
