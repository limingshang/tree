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
	Value  int     // 当前节点值
	Left   *Node   // 当前节点的左节点
	Right  *Node   // 当前节点的右节点
	Height float64 // 当前节点的高度
	Parent *Node   // 当前节点的父节点
}

// 前序遍历
func (n Node) AfterRange() {
	if n != (Node{}) {
		fmt.Println(n.Value)
		n.Left.AfterRange()
		n.Right.AfterRange()
	}
}

// 中序遍历
func (n Node) MiddleRange() {
	if n != (Node{}) {
		n.Left.AfterRange()
		fmt.Println(n.Value)
		n.Right.AfterRange()
	}
}

// 后序遍历
func (n Node) FloorRange() {
	if n != (Node{}) {
		n.Left.AfterRange()
		n.Right.AfterRange()
		fmt.Println(n.Value)
	}
}

// 定义接口
type Tree interface {
	CreateNode(int) Node		// 创建节点方法
	Insert(*Node, int) Node	    // 插入节点方法
	Delete(*Node, int) Node		// 删除节点方法
	Find(*Node, int) Node		// 查找节点方法
}
