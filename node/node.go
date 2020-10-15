// +----------------------------------------------------------------------
// | 功能:
// +----------------------------------------------------------------------
// | Copyright (c) a18203232363@gmail.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://github.com/limingshang )
// +----------------------------------------------------------------------
// | Author: apple 
// +----------------------------------------------------------------------
// | Date:  8:24 下午
// +----------------------------------------------------------------------

package node

import "fmt"


type Node struct {
	value  int     // 当前节点值
	left   *Node   // 当前节点的左节点
	right  *Node   // 当前节点的右节点
	height float64 // 当前节点的高度
	parent *Node   // 当前节点的父节点
}

// 前序遍历
func (n Node) AfterRange() {
	if n != (Node{}) {
		fmt.Println(n.value)
		n.left.AfterRange()
		n.right.AfterRange()
	}
}

// 中序遍历
func (n Node) MiddleRange() {
	if n != (Node{}) {
		n.left.AfterRange()
		fmt.Println(n.value)
		n.right.AfterRange()
	}
}

// 后序遍历
func (n Node) FloorRange() {
	if n != (Node{}) {
		n.left.AfterRange()
		n.right.AfterRange()
		fmt.Println(n.value)
	}
}

// 求当前节点的最大高度
func (n Node) Height() float64 {
	if n == (Node{}) {
		return 0
	}
	leftHeight := n.left.Height()
	rightHeight := n.right.Height()
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}
// 定义接口
type Tree interface {
	CreateNode(int) Node		// 创建节点方法
	Insert(*Node, int) Node	    // 插入节点方法
	Delete(*Node, int) Node		// 删除节点方法
	Find(*Node, int) Node		// 查找节点方法
}
