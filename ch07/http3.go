// Http Interface.
// http.ServeMux 实现路由
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

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for i, v := range db {
		fmt.Fprintf(w, "%s: %s\n", i, v)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	// 获取item参数值.
	item := req.URL.Query().Get("item")
	v, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) //404 必须先写
		fmt.Fprintf(w, "找不到：%q\n", item)
		return
	}
	fmt.Fprintf(w, "价格是：%s\n", v)
}

func main() {
	db := database{"shoes": 50, "sock": 6}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list)) // 不是http.HandleFunc
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8899", mux))
}
