package main

import "fmt"

func solution(n int, edge [][]int) int {
	Nodes := make([][]int, n+1)

	for _, e := range edge {
		u, v := e[0], e[1]
		Nodes[u] = append(Nodes[u], v)
		Nodes[v] = append(Nodes[v], u)
	}

		var depth int
		depthPtr := &depth
		*depthPtr = 1;

		visitied := make([]int, n + 1)
		nodeMaps := make(map[int][]int)
		depthMaps := make(map[int]int)
	for i := 0; i <  len(visitied); i++ {
		visitied[i] = 0
	}
	fmt.Println(Nodes)


	for i := 1; i < len(Nodes); i++ {
		fmt.Println(nodeMaps)
		if visitied[i] != 1{
			visitied[i] = 1
			if len(Nodes[i]) == 0 {
				depthMaps[i] = 1
			 } else {
			depthMaps[i] = depth
			nodeMaps[depth] = append(nodeMaps[depth], i)
		}
	}

		for _, v := range Nodes[i] {
			if visitied[v] != 1{
			*depthPtr = depthMaps[i]  + 1
			visitied[v] = 1
			nodeMaps[depth] = append(nodeMaps[depth], v)
			depthMaps[v] = depth
			}
		}
	}

	maxDepth := 0

	for k, _ := range nodeMaps {
		if maxDepth <  k{
			maxDepth = k
		}	
	}

	count := len(nodeMaps[maxDepth])

	return count
}

func main() {
	test := [][]int{{4, 3}, {1, 3}, {2, 3}}
	n := 4

	fmt.Print(solution(n, test))
}
/*

print(solution(4, [[4, 3], [1, 3], [2, 3]]), 2)

*/