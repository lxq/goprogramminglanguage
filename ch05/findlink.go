// 从标准输入中读取的HTML文档中的所有链接，并输出.
// 调用方式(结合第1章的fetch)：>./fetch.exe http://www.weetgo.com/ | ./findlink.exe
// 5.2 节
// @date 2019/6/2
// @author Allen Lin
// @email xqlin@qq.com
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if nil != err {
		fmt.Fprintf(os.Stderr, "获取HTML失败：%v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit 遍历HTML树上的所有节点，抽取链接元素的href属性值.
// 采用递归调用.
func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := node.FirstChild; nil != c; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
