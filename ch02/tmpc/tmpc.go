/**
 * 包练习.温度转换
 * 2.6 节
 *
 * @date 2019/5/23
 * @author Allen Lin
 * @email xqlin@qq.com
 */
// 包名必须是目录名，其他包的文件统一放到该包下.
package tmpc

import "fmt"

// 摄氏度类型
type Celsuis float64

// 华氏度类型
type Fahrenheit float64

const (
	AbsZeroC Celsuis = -273.15
	FreezC   Celsuis = 0
	BoilC            = 100
)

func (c Celsuis) String() string {
	return fmt.Sprintf("%g℃", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g℉", f)
}
