/**
 * http包构建Server练习.
 *
 * @date 2019/5/22
 * @author Allen Lin
 * @email xqlin@qq.com
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// 统计被访问的次数
var count int

// 互斥量，用于避免并行访问时计数冲突.
var mu sync.Mutex

func main() {
	fmt.Println("simple http server.")

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8899", nil))
}

// http的标准路由处理函数，参数列表固定.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "URL: %q\n", r.URL.Path)
}

// 计数路由
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "首页被访问次数：%d\n", count)
	mu.Unlock()
}
