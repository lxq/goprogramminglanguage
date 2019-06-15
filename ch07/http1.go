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
func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for i, v := range db {
		fmt.Fprintf(w, "%s: %s\n", i, v)
	}
}

func main() {
	db := database{"shoes": 50, "sock": 6}
	log.Fatal(http.ListenAndServe("localhost:8899", db))
}
