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
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if nil != err {
			fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if nil != err {
			fmt.Fprintf(os.Stderr, "读取错误：%v \n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}
