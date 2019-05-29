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
	"fmt"
	"sort"
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

}
