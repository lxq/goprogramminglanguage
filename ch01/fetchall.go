/**
 * 获取多个URL，练习Go 并行.
 *
 * @date 2019/5/22
 * @author Allen Lin
 * @email xqlin@qq.com
 */
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Golang goroutine test.")

	// time包对于时间操作很方便.
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		// start goroutine
		go fetch(url, ch)
	}
	// 下面range只是方便，并无其他特殊意义.
	for range os.Args[1:] {
		// 接收channel数据
		fmt.Println(<-ch)
	}

	fmt.Printf("总耗时: %.2fs\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	res, err := http.Get(url)
	if nil != err {
		// 错误信息发送给channel
		ch <- fmt.Sprint(err)
		return
	}
	// io.Copy()很方便.
	// ioutil.Discard 表示结果被丢弃.
	bytes, err := io.Copy(ioutil.Discard, res.Body)
	// 关闭，以防泄漏
	res.Body.Close()
	if nil != err {
		ch <- fmt.Sprintf("读取错误 %s: %v", url, err)
		return
	}
	n := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs	%7d	%s", n, bytes, url)
}
