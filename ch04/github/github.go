// Package github
// JSON example
// 4.5 节
// @date 2019/6/2
// @author Allen Lin
// @email xqlin@qq.com

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// User github user
type User struct {
	Login string
	// 成员标签定义 P84
	HTMLURL string `json:"html_url"`
}

// Issue 表达了github的issue信息.
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

// SearchResult 是Issue的查寻结果
type SearchResult struct {
	TotalCout int `json:"total_count"`
	Items     []*Issue
}

// IssueURL Gitbug上的接口URL.
const IssueURL = "https://api.github.com/search/issues"

// SearchIssues 查询github的issue跟踪接口.
func SearchIssues(terms []string) (*SearchResult, error) {
	// http get
	q := url.QueryEscape(strings.Join(terms, " "))
	res, err := http.Get(IssueURL + "?q=" + q)
	if nil != err {
		return nil, err
	}
	defer res.Body.Close()

	// get error
	if http.StatusOK != res.StatusCode {
		return nil, fmt.Errorf("查寻错误: $s", res.Status)
	}

	// issue result
	var issRes SearchResult
	// 采用流式解码器
	if err := json.NewDecoder(res.Body).Decode(&issRes); nil != err {
		return nil, err
	}
	return &issRes, nil
}
