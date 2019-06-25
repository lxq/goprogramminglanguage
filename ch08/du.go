// UNIX du 命令实现.
// 8.8 节
// @date 2019/6/25
// @author Allen Lin
// @email xqlin@qq.com

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// 初始目录
	flag.Parse()
	root := flag.Arg(0)
	if "" == root {
		root = "."
	}

	// 遍历目录
	sizes := make(chan int64)
	go func() {
		walkDir(root, sizes)
		close(sizes)
	}()

	// output
	var nfiles, nbytes int64
	for n := range sizes { // range实现channel的接收
		nfiles++
		nbytes += n
	}
	fmt.Printf("共%d个文件，总共大小%.2fMB.", nfiles, float64(nbytes)/1e6)
}

// walkDir 遍历目录树，并把查到的文件的大小发送给sizes
func walkDir(dir string, sizes chan<- int64) {
	for _, entry := range dirlists(dir) {
		if entry.IsDir() {
			sub := filepath.Join(dir, entry.Name())
			walkDir(sub, sizes)
		} else {
			sizes <- entry.Size()
		}
	}
}

// dirlists 获取目录中的所有条目.
func dirlists(dir string) []os.FileInfo {
	lists, err := ioutil.ReadDir(dir)
	if nil != err {
		fmt.Printf("获取目录条目失败：%v\n", err)
		return nil
	}
	return lists
}
