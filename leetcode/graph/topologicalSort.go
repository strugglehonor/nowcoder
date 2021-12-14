package graph

// 拓扑排序，把入度为0的点，打印，然后去除该点及其影响，再找入度为0的点，依次类推
func (graph Graph) topologicalSort() []*Node {
	inMap := map[*Node]int{}
	zeroQueue := make([]*Node, 0)
	for _, node := range graph.nodes {
		inMap[node] = node.in
		if node.in == 0 {
			zeroQueue = append(zeroQueue, node)
		}
	}

	result := make([]*Node, 0)
	for len(zeroQueue) != 0 {
		// 弹出第一个元素
		cur := zeroQueue[0]
		zeroQueue = append(zeroQueue[:0], zeroQueue[1:]...)
		result = append(result, cur)
		// 处理所有下一个节点
		for _, next := range cur.nexts {
			inMap[next]--
			if inMap[next] == 0 {
				zeroQueue = append(zeroQueue, next)
			}
		}
	}

	return result
}