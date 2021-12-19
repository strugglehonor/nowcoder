package graph


func NewGraph(matrix [][]int) Graph {
    graph := Graph{}
    graph.nodes = make(map[int]*Node)
    n := len(matrix)
    var fromNode, toNode *Node
    for i:=0; i<n; i++ {
        // fromNode = &Node{val: matrix[i][0]}
        // toNode = &Node{val: matrix[i][1]}
        from, to := matrix[i][0], matrix[i][1]
        if _, ok := graph.nodes[from]; !ok {
            fromNode = &Node{val: from}
            graph.nodes[from] = fromNode
        }
        if _, ok := graph.nodes[to]; !ok {
            toNode = &Node{val: to}
            graph.nodes[to] = toNode
        }
        // 这句代码不能漏了，否则程序的tonode和from的Out in是不对的。
        fromNode, toNode = graph.nodes[from], graph.nodes[to]
        fromNode.out++
        toNode.in++
    }
    return graph
}

func findJudge(n int, trust [][]int) int {
    if len(trust) < 1 && n==1 {
        return 1
    }
    graph := NewGraph(trust)
    for k, v := range graph.nodes {
        if v.out == 0 && v.in == n-1 {
            return k
        }
    }
    return -1
}