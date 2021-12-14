package graph

// 图 BFS 遍历，使用队列
func (graph Graph) bfs(node *Node) []*Node {
	res := make([]*Node, 0)
	nodeQueue := make([]*Node, 0)
	nodeSet := map[*Node]int{}

	nodeQueue = append(nodeQueue, node)
	nodeSet[node] = node.val

	for len(nodeQueue) != 0 {
		cur := nodeQueue[0]
		nodeQueue = append(nodeQueue[0:], nodeQueue[1:]...)
		res = append(res, cur)
		for _, no := range cur.nexts {
			if _, ok := nodeSet[no]; !ok {
				res = append(res, no)
				nodeSet[no] = no.val
			}
		}
	}
	
	return res
}

// 图 DFS 遍历，使用栈
func (graph Graph) dfs(node *Node) []*Node {
	res := make([]*Node, 0)
	nodeStack := make([]*Node, 0)
	nodeSet := map[*Node]int{}

	nodeStack = append(nodeStack, node)
	nodeSet[node] = node.val
	res = append(res, node)

	for len(nodeStack) != 0 {
		cur := nodeStack[len(nodeStack)-1]
		for _, no := range cur.nexts {
			if _, ok := nodeSet[node]; !ok {
				nodeStack = append(nodeStack, no)
				nodeSet[no] = no.val
				res = append(res, no)
			}
		}
	}
	return res
}