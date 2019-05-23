/**
 * flag包练习指针.
 * 2.3.2 节
 *
 * @date 2019/5/22
 * @author Allen Lin
 * @email xqlin@qq.com
 */

package main

import (
	"flag"
	"fmt"
	"strings"
)

// flag包内容基本都以指针返回.
var n = flag.Bool("n", false, "行尾是否换行")
var s = flag.String("s", ":", "选项分割符")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *s))
	if !*n {
		fmt.Println()
	}
}
