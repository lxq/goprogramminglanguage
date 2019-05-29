/**
 * struct tree example
 * 4.4 节
 *
 * @date 2019/5/29
 * @author Allen Lin
 * @email xqlin@qq.com
 */

package main

import (
	"fmt"
)

// struct成员顺序决定了同一性
type tree struct {
	val int
	// 聚合类型（数组与结构）不能包含自己，但可以包含自己类型的指针(可用于实现链表和树)
	left, right *tree
}

// 往二叉树加入一个元素
func add(t *tree, val int) *tree {
	if nil == t {
		//
		t = new(tree)
		t.val = val
		return t
	}
	if val < t.val {
		// 加入左子树
		t.left = add(t.left, val)
	} else {
		// 递归
		t.right = add(t.right, val)
	}
	return t
}

// 遍历二叉树，把节点值追加到 []int中
// 返回[]int，是为了递归调用.
func appendVal(vals []int, t *tree) []int {
	if nil != t {
		// 先左子树
		vals = appendVal(vals, t.left)
		vals = append(vals, t.val)
		vals = appendVal(vals, t.right)
	}
	return vals
}

// Sort 对整数进行排序
func Sort(vals []int) {
	// 先生成二叉树
	var root *tree
	for _, v := range vals {
		root = add(root, v)
	}
	// 把二叉树结果导出成slice
	appendVal(vals[:0], root)
}

func main() {
	fmt.Println("自定义二叉树练习.")
	s := []int{123, 3, 345, 0, 13, 789}
	fmt.Printf("原始数组：%v\n", s)
	Sort(s)
	fmt.Printf("二叉树排序后：%v\n", s)
}
