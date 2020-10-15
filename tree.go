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

import "tree/balance"


func main()  {
	balanceTree()
}

func balanceTree()  {
	node:= balance.CreateNode(10)
	balance.InsertNode(&node, 50)
	balance.InsertNode(&node, 60)
	balance.InsertNode(&node, 30)
	balance.InsertNode(&node, 70)
	balance.InsertNode(&node, 20)
	balance.InsertNode(&node, 40)
	balance.InsertNode(&node, 120)
	balance.MiddleRange(node)
}


