// TCP服务器.
// 8.2 节
// @date 2019/6/20
// @author Allen Lin
// @email xqlin@qq.com

package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// "tcp"不能大写，必须全部小写.
	listener, err := net.Listen("tcp", "localhost:8899")
	if nil != err {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if nil != err {
			log.Println(err)
			continue
		}

		go handlConn(conn)
	}
}

// 处理Client连接.
func handlConn(c net.Conn) {
	defer c.Close()

	for {
		// 写字符串使用io.WriteString()效率更高.
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if nil != err {
			return // 如，连接断开
		}
		time.Sleep(1 * time.Second)
	}
}
