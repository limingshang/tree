// +----------------------------------------------------------------------
// | 功能: 各种树调用
// +----------------------------------------------------------------------
// | Copyright (c) a18203232363@gmail.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://github.com/limingshang )
// +----------------------------------------------------------------------
// | Author: apple 
// +----------------------------------------------------------------------
// | Date:  5:09 下午
// +----------------------------------------------------------------------

package main

import "github.com/limingshang/tree/preface"

func main() {
	prefaceTree()
}

func prefaceTree()  {
	// 普通树插入因为不涉及各种计算，直接对比下一个节点的大小做插入，所以最后执行的结果为
	//              20
	//        10       30
	//     4               50
	// 3      7         40    60
	//      6   8
	//    5

	var node preface.Node
	preface.Insert(&node, 20)
	preface.Insert(&node, 10)
	preface.Insert(&node, 30)
	preface.Insert(&node, 50)
	preface.Insert(&node, 60)
	preface.Insert(&node, 40)
	preface.Insert(&node, 4)
	preface.Insert(&node, 3)
	preface.Insert(&node, 7)
	preface.Insert(&node, 8)
	preface.Insert(&node, 6)
	preface.Insert(&node, 5)
	preface.AfterRange(node)
	//preface.MiddleRange(node)
	//fmt.Println(node)
}