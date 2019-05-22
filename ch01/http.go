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
)

func main() {
	fmt.Println("simple http server.")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8899", nil))
}

// http的标准路由处理函数，参数列表固定.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL: %q\n", r.URL.Path)
}
