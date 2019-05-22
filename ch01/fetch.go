/**
 * URL 练习.
 *
 * @date 2019/5/121
 * @author Allen Lin
 * @email xqlin@qq.com
 */

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// http.Get()产生一个HTTP请求.
		res, err := http.Get(url)
		if nil != err {
			fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
			os.Exit(1)
		}
		// ReadAll()一次读取全部内容.
		// body, err := ioutil.ReadAll(res.Body)
		// TODO: 下面不能用: _, err := ，短赋值符需要有至少一个新的变量，所以改成：_, e :=
		_, e := io.Copy(os.Stdout, res.Body)
		res.Body.Close()
		if nil != e {
			fmt.Fprintf(os.Stderr, "读取错误：%v \n", e)
			os.Exit(1)
		}
		// fmt.Printf("%s", body)
	}
}
