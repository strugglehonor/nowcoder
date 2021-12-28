package graph

func canFinish(numCourses int, prerequisites [][]int) bool {
    n := len(prerequisites)
    if n == 0 {
        return true
    }
    // nodeMap := make(map[int]Node)
    var fromNode, toNode *Node
    graph := Graph{}
    // graph.nodes = map[int]&Node{}
    graph.nodes = map[int]*Node{}
    // 建立图
    for i:=0; i<n; i++ {
        fromval, toval := prerequisites[i][1], prerequisites[i][0]
        // 如果节点不存在，新建节点
        if _, ok := graph.nodes[fromval]; !ok {
            fromNode = &Node{val: fromval}
            graph.nodes[fromval] = fromNode
        }
        if _, ok := graph.nodes[toval]; !ok {
            toNode = &Node{val: toval}
            graph.nodes[toval] = toNode
        }
        fromNode, toNode = graph.nodes[fromval], graph.nodes[toval]

        fromNode.out++
        toNode.in++
        fromNode.nexts = append(fromNode.nexts, toNode)  // 为什么这行代码和位置有关
        edge := &Edge{
            to: toNode,
            from: fromNode,
            weight: 0,
        }
        graph.edges = append(graph.edges, edge)
        fromNode.edges = append(fromNode.edges, edge)
        toNode.edges = append(toNode.edges, edge)
        graph.nodes[fromval], graph.nodes[toval] = fromNode, toNode

    }
    nodeRes := topologicalSort(graph)
    for nodeval := range graph.nodes {
        res := compare(nodeval, nodeRes)
        if !res {
            return false
        }
    }
    return true
}

// 数组中是否存在元素x，返回bool
func compare(x int, nums []int) bool {
    for _, num := range nums {
        if num == x {
            return true
        }
    }
    return false
}

// 拓扑排序，把入度为0的点，打印，然后去除该点及其影响，再找入度为0的点，依次类推
func topologicalSort(graph Graph) []int {
	inMap := map[int]int{}  // 存储Node的入度的map
	zeroQueue := make([]*Node, 0)  // 存储入度为0的node的队列

	for _, node := range graph.nodes {
		inMap[node.val] = node.in
		if node.in == 0 {
			zeroQueue = append(zeroQueue, node)
		}
	}

	result := make([]int, 0)
	for len(zeroQueue) != 0 {
		// 弹出第一个元素
		cur := zeroQueue[0]
		zeroQueue = append(zeroQueue[:0], zeroQueue[1:]...)
		result = append(result, cur.val)
		// 处理所有下一个节点
		for _, next := range cur.nexts {
			inMap[next.val]--
            // fmt.Println(inMap[next.val])
			if inMap[next.val] == 0 {
				zeroQueue = append(zeroQueue, next)
			}
		}
	}
	return result
}