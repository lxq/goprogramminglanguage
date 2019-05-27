/**
 * cypto/sha256包练习数组比较
 * 4.1 节
 *
 * @date 2019/5/26
 * @author Allen Lin
 * @email xqlin@qq.com
 */

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println("sh256:数组比较")
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	// %x ：十六进制， %t:布尔值，%T：值类型
	fmt.Printf("sha256 c1:%x\nsha256 c2:%x\nc1==c2:%t\n%T\n", c1, c2, c1 == c2, c1)
}
