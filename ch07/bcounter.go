// Interface.
// 7.1 节
// @date 2019/6/13
// @author Allen Lin
// @email xqlin@qq.com

package main

import (
	"fmt"
)

// Bcounter 自定义类型
type Bcounter int

// Write 实现io.Writer接口.
func (c *Bcounter) Write(p []byte) (int, error) {
	*c += Bcounter(len(p))
	return len(p), nil
}

func main() {
	fmt.Println("自定义Interface.")

	var c Bcounter
	c.Write([]byte("Hello."))
	fmt.Println(c)

	c = 0
	name := "weetgo.com"
	// Bcounter实现了io.Writer接口
	fmt.Fprintf(&c, "%s", name)
	fmt.Println(c)
}
