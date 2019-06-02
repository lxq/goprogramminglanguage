/**
 * struct匿名成员练习
 * 4.4.3 节
 *
 * @date 2019/5/29
 * @author Allen Lin
 * @email xqlin@qq.com
 */

package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point  //Center
	Radius int
}

type Wheel struct {
	Circle     //
	Spokes int // 条辐
}

func main() {
	fmt.Println("Struct匿名成员.")
	// 匿名成员初始化
	w1 := Wheel{Circle{Point{8, 8}, 5}, 20}
	w2 := Wheel{
		Circle: Circle{
			Point: Point{11, 12},
			// TODO: 下面一行的逗号和spokes后面的逗号不能省略.目前还不明白为什么！！！——原因是用于区分下一行，如果把}放同一行就可以不需要逗号.
			Radius: 150,
		},
		Spokes: 21,
	}
	fmt.Printf("%#v\n", w1)
	fmt.Printf("%#v\n", w2)
}
