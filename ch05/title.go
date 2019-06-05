// Go 函数练习.
// 5.8 节
// @date 2019/6/5
// @author Allen Lin
// @email xqlin@qq.com

package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// >./title.exe http://www.weetgo.com
func main() {
	fmt.Println("延迟调用.")

	url := os.Args[1]
	title(url)
}

func title(url string) error {
	resp, err := http.Get(url)
	if nil != err {
		return err
	}
	defer resp.Body.Close()

	// 检查 Content-Type
	ct := resp.Header.Get("Content-Type")
	if "text/html" != ct && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s 文档类型是：%s，非text/html.", url, ct)
	}
	// x/net/http => http.Parse()
	doc, err := html.Parse(resp.Body)
	if nil != err {
		return fmt.Errorf("把%s解析成HTML出错：%v.", url, err)
	}

	// 匿名函数，也叫闭包.
	visit := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEach(doc, visit, nil)
	return nil
}

// 遍历html节点树，	子节点被访问前调用pre()，访问后调用post().
// 函数参数
func forEach(n *html.Node, pre, post func(n *html.Node)) {
	if nil != pre {
		pre(n)
	}

	for c := n.FirstChild; nil != c; c = c.NextSibling {
		// 递归
		forEach(c, pre, post)
	}

	if nil != post {
		post(n)
	}
}
