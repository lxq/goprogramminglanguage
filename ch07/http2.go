// Http Interface.
// 7.7 节
// @date 2019/6/16
// @author Allen Lin
// @email xqlin@qq.com

package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

// 实现String()方法供后面格式时自动调用.
func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// implemetn http.Handler interface.
// 根据参数选择处理.
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for i, v := range db {
			fmt.Fprintf(w, "%s: %s\n", i, v)
		}
	case "/price":
		// 获取item参数值.
		item := req.URL.Query().Get("item")
		v, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) //404 必须先写
			fmt.Fprintf(w, "找不到：%q\n", item)
			return
		}
		fmt.Fprintf(w, "价格是：%s\n", v)
	default:
		// 另一种返回方式
		msg := fmt.Sprintf("找不到你的网页：%s\n", req.URL)
		http.Error(w, msg, http.StatusNotFound)
	}
}

func main() {
	db := database{"shoes": 50, "sock": 6}
	log.Fatal(http.ListenAndServe("localhost:8899", db))
}
