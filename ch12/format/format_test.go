// format包测试.
// @date 2019/6/26
// @author Allen Lin
// @email xqlin@qq.com
package format

import (
	"testing"
)

func TestBool(t *testing.T) {
	b := true
	if "true" != AnyStr(b) {
		t.Error(`b=true => false`)
	}
	n := -124
	if "false" != AnyStr(n) {
		t.Error(`-124 is not bool.`)
	}
}

func TestString(t *testing.T) {
	if `"weetgo"` != AnyStr("weetgo") {
		// fmt.Println(AnyStr("weetgo"))
		t.Error(`Type of "weetgo" if not string`)
	}
}
