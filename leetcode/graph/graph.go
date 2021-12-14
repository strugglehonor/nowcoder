package graph

// node，图结构的点
type Node struct {
	in  int  // 入度
	out int  // 出度
	val int  // node的值，类似于编号
	nexts []*Node  // 下个节点
	edges  []*Edge  // 边
}

// edge，图结构的边
type Edge struct {
	weight  int  // 权重
	from  *Node  // 边的from点
	to  *Node  // 边的to点
}

// 图
type Graph struct {
	edges []*Edge
	nodes map[int]*Node
}

// new node
func NewNode(val int) *Node {
	return &Node{
		val: val,
	}
}

// new edge
func NewEdge(weight int, from, to *Node) Edge {
	return Edge{
		weight: weight,
		from: from,
		to: to,
	}
}

// new graph
func newGraph() Graph {
	return Graph{}
}

// [[weight1, from1, to1], [weight2, from2, to2]]（题目中常出现的结果转化为我们的Graph图结构
func GraphGenerator(matrix [][]int) Graph {
	graph := newGraph()
	var fromNode, toNode *Node
	n := len(matrix)
	for i:=0; i<n; i++ {
		weight := matrix[i][0]
		from := matrix[i][1]
		to := matrix[i][2]
		// 如果节点不存在，新建节点，加入图中
		if _, ok := graph.nodes[from]; !ok {
			fromNode = NewNode(from)
			graph.nodes[from] = fromNode
		}
		if _, ok := graph.nodes[to]; !ok {
			toNode = NewNode(to)
			graph.nodes[to] = toNode
		}
		edge := NewEdge(weight, fromNode, toNode)  // init edge
		// init fromNode
		fromNode.nexts = append(fromNode.nexts, toNode)
		fromNode.out++
		fromNode.edges = append(fromNode.edges, &edge)
		// init foNode
		toNode.in++
		toNode.edges = append(toNode.edges, &edge)
	}
	return graph
}