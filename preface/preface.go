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

type Node struct {
	value int   // 当前节点值
	left  *Node // 当前节点的左节点
	right *Node // 当前节点的右节点
}

func createNode(value int) Node {
	return Node{
		value: value,
		left:  &Node{},
		right: &Node{},
	}
}
// 此是插入操作，删除很简单，左节点的最右节点上移活着右节点最左节点上移到要删除的节点就可以了
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


