package graph

func allPathsSourceTarget(graph [][]int) [][]int {
    print(1)
    // var node Node
    n := len(graph)
    nodeMap := make(map[int]*Node)
    for i:=0; i<n; i++ {
        node := getNode(nodeMap, i)
        m := len(graph[i])
        for j:=0; j<m; j++ {
            node.out++
            to := getNode(nodeMap, graph[i][j])
            to.in++
            node.nexts = append(node.nexts, to)
        }
    }
    res, ans := make([]int, 0), make([][]int, 0)
    dfs(graph, nodeMap[0], res, ans)
    print(res)

    return ans
}

func dfs(graph [][]int, node *Node, res []int, ans [][]int) {
    n := len(graph)
    if node.val == n-1 {
        ans = append(ans, res)
        return
    }
    res = append(res, node.val)
    for _, next := range node.nexts {
        dfs(graph, next, res, ans)
    }
}    

func getNode(nodeMap map[int]*Node, val int) *Node {
    var node *Node
    if _, ok := nodeMap[val]; ok {
        node = nodeMap[val]
    }else {
        node = &Node{val: val}
        nodeMap[val] = node
    }
    return node
}

