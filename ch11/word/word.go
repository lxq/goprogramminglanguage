// 用于测试练习的包.
// 11.2 节
// @date 2019/6/26
// @author Allen Lin
// @email xqlin@qq.com
package word

// IsPalindrome 判断是否为对称字符串.
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
