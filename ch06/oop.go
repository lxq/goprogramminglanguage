// 面向对象：方法.
// 6.1 节
// @date 2019/6/9
// @author Allen Lin
// @email xqlin@qq.com
package main

import (
	"fmt"
	"gopl/ch06/geom"
	"image/color"
)

func main() {
	fmt.Println("测试面对对象.")
	p := geom.Point{1, 2}
	q := geom.Point{4, 6}
	fmt.Println(geom.Dist(p, q))
	fmt.Println(p.Dist(q))

	path := geom.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(path.Dist())

	r := &p
	r.Scale(2)
	fmt.Println(*r)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	cp := geom.CPoint{geom.Point{1, 1}, red}
	cp2 := geom.CPoint{geom.Point{4, 4}, blue}
	fmt.Println(cp)
	cp.Scale(3)
	fmt.Println(cp)
	// 下面不能直接传递cp，因为结构内嵌不是继承，而是组合.
	fmt.Println(cp2.Dist(cp.Point))
}
