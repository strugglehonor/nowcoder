package graph

import (
	"math"
	"reflect"
)

// 单源最短路径算法
func (graph Graph) dijkstra(from Node) map[*Node]int {
	// 标记node是否被处理过，key为node的编号，值为node
	nodeMap := map[int]Node{} 
	// from到node的距离，初始化from到from的距离为0（到自己的距离) 
	distanceMap := map[*Node]int{&from: 0} 
	minNode := getMiniDistanceAndUnselectedNode(nodeMap, distanceMap)
	defaultNode := Node{}
	// 如果minNode返回的是一个Node{}，循环退出
	for !reflect.DeepEqual(defaultNode, minNode) {
		for _, edge := range minNode.edge {
			if _, ok := distanceMap[&edge.to]; !ok {
				// 如果没有记录过这个点的路径，那么直接写入distanceMap
				nodeMap[edge.to.val] = edge.to
				distanceMap[&edge.to] = distanceMap[&minNode] + edge.weight
			}else {
				// 如果原来已经记录过到这个点的路径，那么需要选择较小的
				distanceMap[&edge.to] = min(distanceMap[&minNode] + edge.weight, distanceMap[&edge.to])
			}
		}
		minNode = getMiniDistanceAndUnselectedNode(nodeMap, distanceMap)
		// 表示访问过了这个node
		nodeMap[minNode.val] = minNode
	}

	return distanceMap
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 返回最短的路径的点，在distanceMap中最短路径的点，且不在nodeMap中的.
func getMiniDistanceAndUnselectedNode(nodeMap map[int]Node, distanceMap map[*Node]int) Node {
	distance := math.MaxInt32
	minNode := Node{}
	for k, v := range distanceMap {
		if _, ok := nodeMap[k.val]; !ok && v < distance {
			distance = v
			minNode = *k
		}
	}
	return minNode
}