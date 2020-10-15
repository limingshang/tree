// +----------------------------------------------------------------------
// | 功能: 平衡二叉树
// +----------------------------------------------------------------------
// | Copyright (c) a18203232363@gmail.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://github.com/limingshang )
// +----------------------------------------------------------------------
// | Author: apple 
// +----------------------------------------------------------------------
// | Date:  5:50 下午
// +----------------------------------------------------------------------

package balance

import (
	"fmt"
	"github.com/limingshang/tree/node"
	"math"
)

//type Node struct {
//	value  int     // 当前节点值
//	left   *Node   // 当前节点的左节点
//	right  *Node   // 当前节点的右节点
//	height float64 // 当前节点的高度
//}

// 初始化创建根节点
func CreateNode(value int) node.Node {
	return node.Node{
		value:  value,
		left:   &Node{},
		right:  &Node{},
		height: 0,
	}
}

// 右右型数据左旋转操作此操作做了一次左旋转之后，返回当前数据值，重新赋值给父节点
func rrTurnLeft(node Node) Node {
	x := node.right
	node.right = x.left
	x.left = &node
	return *x
}

// 左左型数据左旋转操作此操作做了一次右旋转之后，返回当前数据值，重新赋值给父节点
func llTrunRight(node Node) Node {
	x := node.left
	node.left = x.right
	x.right = &node
	x.height = getHeight(*x)
	return *x
}

// 获取当前节点的高度
func getHeight(node Node) float64 {
	if node == (Node{}) {
		return 0
	}
	leftHeight := getHeight(*node.left)
	rightHeight := getHeight(*node.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

// 平衡二叉树的插入操作
func InsertNode(node *Node, value int) {
	if *node == (Node{}) {
		// 检测如果当前节点是空创建一个节点值返回
		*node = CreateNode(value)
	} else if value > node.value {
		InsertNode(node.right, value)
		height := getHeight(*node.right) - getHeight(*node.left)
		if math.Abs(height) >= 2 {
			// 当前插入数据大于右侧节点的值为右右插入RR型，需要进行一次左旋转
			if value > node.right.value {
				*node = rrTurnLeft(*node)
			// // 右侧插入 RL 型先进行一次右旋转，在进行一次左旋转
			} else if value < node.right.value {
				*node.right = llTrunRight(*node.right)
				*node = rrTurnLeft(*node)
			}
		}
	} else if value < node.value {
		InsertNode(node.left, value)
		height := getHeight(*node.right) - getHeight(*node.left)
		if math.Abs(height) >= 2 {
			// 当前插入数据大于右侧节点的值为右右插入RR型，需要进行一次左旋转
			if value < node.left.value {
				*node = llTrunRight(*node)
				// 当前插入数据小于右侧节点的值为右右插入RL型，需要现进行一次右旋转然后左旋转一次
			} else if value > node.left.value {
				*node.left = rrTurnLeft(*node)
				*node = llTrunRight(*node)
			}
		}
	} else {

	}
	fmt.Println(*node)
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
