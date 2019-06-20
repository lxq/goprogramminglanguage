// TCP客户端.
// 8.2 节
// @date 2019/6/20
// @author Allen Lin
// @email xqlin@qq.com

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// "tcp"不能大写，必须全部小写.
	conn, err := net.Dial("tcp", "localhost:8899")
	if nil != err {
		log.Fatal(err)
	}
	defer conn.Close()

	netCopy(os.Stdout, conn)
}

// copy网络数据到输出
func netCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); nil != err {
		log.Fatal(err)
	}
}
