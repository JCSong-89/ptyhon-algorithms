package main

type Node struct {
	Data int
	Neighbor *[]Node
}

func dfs(node Node){
	discoverd := make(map[int]int)
	stack := make([]Node, 0)

	stack = append(stack, node)
	discoverd[node.Data] = 1
	stackPtr := &stack

 for; len(stack) != 0; {
	 next := pop(stackPtr)

	 for _, v := range *next.Neighbor {
		 if(discoverd[v.Data] == 0) {
			stack = append(*stackPtr, v)
			discoverd[v.Data] = 1 
		 }
	 }
 }
}

func pop(nodes *[]Node) *Node {
	popNode := (*nodes)[len(*nodes) - 1]
	*nodes = (*nodes)[:len(*nodes) - 1]

	return &popNode
}