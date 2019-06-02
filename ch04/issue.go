// github模块的JSON练习.
// 测试命令 :>.\issue.exe repo:golang/go is:open json decoder
package main

import (
	"fmt"
	"gopl/ch04/github"
	"log"
	"os"
)

func main() {
	res, err := github.SearchIssues(os.Args[1:])
	if nil != err {
		log.Fatal(err)
	}

	fmt.Printf("issues数量:%d\n", res.TotalCout)
	for _, item := range res.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
