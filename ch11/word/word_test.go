// 用于测试练习的包.
// 11.2 节
// @date 2019/6/26
// @author Allen Lin
// @email xqlin@qq.com
package word

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated")==false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak")=false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome")=true`)
	}
}
