// 面向对象：方法.
// 6.1 节
// @date 2019/6/9
// @author Allen Lin
// @email xqlin@qq.com

package geom

import (
	"image/color"
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

// Scale 进行标量缩放.
func (p *Point) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

// Path 直线段
type Path []Point

// Dist 直线段Path的距离.
func (p Path) Dist() float64 {
	d := 0.0
	for i := range p {
		if i > 0 {
			d += p[i-1].Dist(p[i])
		}
	}
	return d
}

// CPoint Point with color.
type CPoint struct {
	Point
	Color color.RGBA
}
