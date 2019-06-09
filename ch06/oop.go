// 面向对象：方法.
// 6.1 节
// @date 2019/6/9
// @author Allen Lin
// @email xqlin@qq.com
package main

import (
	"fmt"
	"gopl/ch06/geom"
)

func main() {
	fmt.Println("测试面对对象.")
	p := geom.Point{1, 2}
	q := geom.Point{4, 6}
	fmt.Println(geom.Dist(p, q))
	fmt.Println(p.Dist(q))
}
