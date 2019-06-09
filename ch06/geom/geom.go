// 面向对象：方法.
// 6.1 节
// @date 2019/6/9
// @author Allen Lin
// @email xqlin@qq.com

package geom

import (
	"math"
)

// Point 2D点.
type Point struct {
	X, Y float64
}

// Dist 常规函数计算2点间距离
func Dist(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

// Dist 计算2个Point间的距离.
func (p Point) Dist(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}
