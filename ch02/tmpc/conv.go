/**
 * 包练习.温度转换
 * 2.6 节
 *
 * @date 2019/5/23
 * @author Allen Lin
 * @email xqlin@qq.com
 */
package tmpc

// 摄氏转华氏
func CtoF(c Celsuis) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// 华氏转摄氏
func FtoC(f Fahrenheit) Celsuis {
	return Celsuis((f - 32) * 5 / 9)
}
