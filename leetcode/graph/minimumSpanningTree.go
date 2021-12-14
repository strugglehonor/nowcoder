package graph

// 最小生成树算法：
// 算法使得图的任意两个节点都是连通的，且边的总权值要最小
// 两个算法，K算法和P算法

// P算法,node为从这个点出发
func (graph Graph) Prim(node *Node) []*Edge {
	res := make([]*Edge, 0)
	// node set 是为了判断，防止重复
	nodeSet := map[int]*Node{node.val: node}
	// 这里应该使用小根堆效率较高，为了简单，使用了切片
	edgeHeap := make([]*Edge, 0)  
	// 从这个点出发的所有的边加到小根堆
	for _, edge := range node.edges {
		edgeHeap = append(edgeHeap, edge)
	}
	for len(edgeHeap) != 0 {
		edge := minEdge(edgeHeap)  // 权值最小的边（如果是小根堆就可以直接返回)
		nodeval := edge.to.val
		// 当前这个点不在set中，这条边要加入，并且把这个点加入到set
		if _, ok := nodeSet[nodeval]; !ok {
			nodeSet[nodeval] = edge.to
			res = append(res, edge)
			// 把这个点关联的所有边加到小根堆
			for _, edge := range edge.to.edges {
				edgeHeap = append(edgeHeap, edge)
			}
		}

	}
	return res
}

// K算法，
func (graph Graph) Kruskal() []*Edge {
	nodesMap := generateNodeSet(graph)
	edgeRes := make([]*Edge, 0)  // 返回需要的边

	// 遍历所有的边，先从小的边开始
	for i:=0; i<len(graph.edges); i++ {
		edge := minEdge(graph.edges)
		node1, node2 := edge.from, edge.to
		node1Set, node2Set := nodesMap[node1.val], nodesMap[node2.val]

		isSame := isSameNodeSet(node1Set, node2Set)
		if !isSame {
			edgeRes = append(edgeRes, edge)
			unionNodeSet(*node1Set, *node2Set)
		}
	}

	return edgeRes
}

// 返回权值最小的边
func minEdge(edges []*Edge) *Edge {
	res := edges[0]
	for _, edge := range edges {
		if edge.weight > res.weight {
			res = edge
		}
	}
	return res
}

// 将图中所有的点加入分成集合 []map[int]*[]Node
func generateNodeSet(graph Graph) map[int]*[]*Node {
	// res := make([]map[int]*[]Node, 0)
	var nodesMap map[int]*[]*Node
	for _, node := range graph.nodes {
		nodeArray := make([]*Node, 0)
		nodeArray = append(nodeArray, node)
		nodesMap[node.val] = &nodeArray
		// res = append(res, nodesMap)
	}
	return nodesMap
}

// 将两个集合中的所有元素合并成一个大集合,ns为nodeset的意思
func unionNodeSet(ns1, ns2 []*Node) []*Node {
	ns2 = append(ns2, ns1...)
	return ns2
}

// 两个集合是否在一个集合里面
func isSameNodeSet(ns1, ns2 *[]*Node) bool {
	if ns1 == ns2 {
		return true
	}
	return false
}