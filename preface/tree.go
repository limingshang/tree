// +----------------------------------------------------------------------
// | 功能: 普通二叉树实现
// +----------------------------------------------------------------------
// | Copyright (c) a18203232363@gmail.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://github.com/limingshang )
// +----------------------------------------------------------------------
// | Author: apple 
// +----------------------------------------------------------------------
// | Date:  11:29 上午
// +----------------------------------------------------------------------

package preface

import "fmt"

type Node struct {
	value  int   // 当前节点值
	left   *Node // 当前节点的左节点
	right  *Node // 当前节点的右节点
	height int   // 当前节点所在的高度
}

func createNode(value int) Node {
	return Node{
		value:  value,
		left:   &Node{},
		right:  &Node{},
		height: value,
	}
}
func Insert(node *Node, value int) {
	if *node == (Node{}) {
		*node = createNode(value)
	} else {
		if value > node.value {
			if *node.right == (Node{}) {
				*node.right = createNode(value)
			} else {
				Insert(node.right, value)
			}
		} else if value < node.value {
			if *node.left == (Node{}) {
				*node.left = createNode(value)
			} else {
				//*node.left = createNode(value)
				Insert(node.left, value)
			}
		}
	}
}

func AfterRange(node Node) {
	if node != (Node{}) {
		fmt.Println(node.value)
		AfterRange(*node.left)
		AfterRange(*node.right)
	}
}
func MiddleRange(node Node) {
	if node != (Node{}) {
		MiddleRange(*node.left)
		fmt.Println(node.value)
		MiddleRange(*node.right)
	}
}
func FloorRange(node Node) {
	if node != (Node{}) {
		FloorRange(*node.left)
		FloorRange(*node.right)
		fmt.Println(node.value)
	}
}
