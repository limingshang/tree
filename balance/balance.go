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
	"math"
)

type Node struct {
	value  int   // 当前节点值
	left   *Node // 当前节点的左节点
	right  *Node // 当前节点的右节点
	height float64   // 当前节点的高度
}

func CreateNode(value int) Node {
	return Node{
		value:  value,
		left:   &Node{},
		right:  &Node{},
		height: 0,
	}
}

// 求当前节点的最大高度
func Height(node *Node) float64 {
	if *node == (Node{}) {
		return 0
	}
	leftHeight := Height(node.left)
	rightHeight := Height(node.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

// 左旋转
func leftRelve(node *Node) Node {
	x := *node.right
	*node.right = *x.left
	*x.left = *node
	x.height = Height(&x)
	node.height = Height(node)
	return x
}

// 右旋转
func RightRelve(node *Node) Node {
	x := *node.left
	*node.left = *x.right
	*x.right = *node
	x.height = math.Max(Height(x.left), Height(x.right))
	node.height = math.Max(Height(node.left), Height(node.right))
	return x
}

// 插入节点
func InsertNode(node *Node, value int) {
	if *node == (Node{}) {
		*node = CreateNode(value)
	} else if value > node.value {
		InsertNode(node.right, value)
		// 出现失衡状态，由于是右侧录入节点，顾直接做旋转
		relve := Height(node.right) - Height(node.left)
		if math.Abs(float64(relve)) == 2 { // 等于2 为右插入不平衡
			// 判断左插入还是右插入
			if value > node.right.value {
				// 右侧插入 RR 型直接进行一次左旋转就好了
				*node = leftRelve(node)
			} else if value < node.right.value {
				// 右侧插入 RL 型先进行一次右旋转，在进行一次左旋转
				*node.right = RightRelve(node.right)
				*node = leftRelve(node)
			}
		}
	} else if value < node.value {
		InsertNode(node.left, value)
		// 出现失衡状态，由于是左侧录入节点，顾直接做旋转
		relve := Height(node.right) - Height(node.left)
		if math.Abs(float64(relve)) == 2 { // 等于2 为右插入不平衡
			// 判断左插入还是右插入
			if value> node.right.value {
				// 右侧插入 LR 型先进行一次左旋转，在进行一次右旋转
				cc := leftRelve(node.left)

				*node.left = cc
				*node = RightRelve(node)

			} else if value < node.right.value {
				// 左侧插入 LL 型直接进行一次左旋转就好了
				*node = RightRelve(node)
			}
		}
	}
	node.height = math.Max(Height(node.left), Height(node.right))
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
