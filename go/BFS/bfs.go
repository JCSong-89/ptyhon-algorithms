package main

type Node struct {
	Data int
	Neighbor *[]Node
}

func bfs(node Node){
	discoverd := make(map[int]int)
	q := make([]Node, 0)

	q = append(q, node)
	discoverd[node.Data] = 1
	qPtr := &q

 for; len(q) != 0; {
	 next := remove(qPtr)

	 for _, v := range *next.Neighbor {
		 if(discoverd[v.Data] == 0) {
			q = append(*qPtr, v)
			discoverd[v.Data] = 1 
		 }
	 }
 }
}

func remove(nodes *[]Node) *Node {
	removeNode := (*nodes)[0]
	*nodes = (*nodes)[1:]

	return &removeNode
}