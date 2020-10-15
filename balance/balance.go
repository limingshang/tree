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
	"math"
	"tree/BaseNode"
)

//type Node struct {
//	Value  int     // 当前节点值
//	Left   *Node   // 当前节点的左节点
//	Right  *Node   // 当前节点的右节点
//	height float64 // 当前节点的高度
//}

// 初始化创建根节点
func CreateNode(Value int) BaseNode.Node {
	return BaseNode.Node{
		Value:  Value,
		Left:   &BaseNode.Node{},
		Right:  &BaseNode.Node{},
		Height: 0,
	}
}

// 右右型数据左旋转操作此操作做了一次左旋转之后，返回当前数据值，重新赋值给父节点
func rrTurnLeft(node BaseNode.Node) BaseNode.Node {
	x := node.Right
	node.Right = x.Left
	x.Left = &node
	return *x
}

// 左左型数据左旋转操作此操作做了一次右旋转之后，返回当前数据值，重新赋值给父节点
func llTrunRight(node BaseNode.Node) BaseNode.Node {
	x := node.Left
	node.Left = x.Right
	x.Right = &node
	x.Height = getHeight(*x)
	return *x
}

// 获取当前节点的高度
func getHeight(node BaseNode.Node) float64 {
	if node == (BaseNode.Node{}) {
		return 0
	}
	leftHeight := getHeight(*node.Left)
	rightHeight := getHeight(*node.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

// 平衡二叉树的插入操作
func InsertNode(node *BaseNode.Node, Value int) {
	if *node == (BaseNode.Node{}) {
		// 检测如果当前节点是空创建一个节点值返回
		*node = CreateNode(Value)
	} else if Value > node.Value {
		InsertNode(node.Right, Value)
		height := getHeight(*node.Right) - getHeight(*node.Left)
		if math.Abs(height) >= 2 {
			// 当前插入数据大于右侧节点的值为右右插入RR型，需要进行一次左旋转
			if Value > node.Right.Value {
				*node = rrTurnLeft(*node)
			// // 右侧插入 RL 型先进行一次右旋转，在进行一次左旋转
			} else if Value < node.Right.Value {
				*node.Right = llTrunRight(*node.Right)
				*node = rrTurnLeft(*node)
			}
		}
	} else if Value < node.Value {
		InsertNode(node.Left, Value)
		height := getHeight(*node.Right) - getHeight(*node.Left)
		if math.Abs(height) >= 2 {
			// 当前插入数据大于右侧节点的值为右右插入RR型，需要进行一次左旋转
			if Value < node.Left.Value {
				*node = llTrunRight(*node)
				// 当前插入数据小于右侧节点的值为右右插入RL型，需要现进行一次右旋转然后左旋转一次
			} else if Value > node.Left.Value {
				*node.Left = rrTurnLeft(*node)
				*node = llTrunRight(*node)
			}
		}
	// 否则是等于当前数据重复插入，不做保存
	} else {
	}
}
//
//func AfterRange(node BaseNode.Node) {
//	if node != node {
//		fmt.Println(node.Value)
//		AfterRange(*node.Left)
//		AfterRange(*node.Right)
//	}
//}
//func MiddleRange(node BaseNode.Node) {
//	if node != (BaseNode.Node{}) {
//		MiddleRange(*node.Left)
//		fmt.Println(node.Value)
//		MiddleRange(*node.Right)
//	}
//}
//func FloorRange(node BaseNode.Node) {
//	if node != (BaseNode.Node{}) {
//		FloorRange(*node.Left)
//		FloorRange(*node.Right)
//		fmt.Println(node.Value)
//	}
//}
