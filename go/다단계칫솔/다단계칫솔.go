package main

import "fmt"

type Node struct {
	Name      string
	Next      []*Node
	Prev *Node
	MyMoney   int // 나의 소지 금액
}

func solution(enroll []string, referral []string, seller []string, amount []int) []int {
	n := len(enroll)
	nodes := make([]Node, 0)
	root := Node{Name: "-",}
	nodes = append(nodes, root)

	for i := 0; i < n; i++ {
		initNode := Node{Name: enroll[i]}
		nodes = append(nodes, initNode)
	} // 각 노드 생성

		 for i := 0; i < n; i++{
			for j := 0; j < n + 1; j++ {
				if referral[i] == nodes[j].Name {
					nodes[j].Next = append(nodes[j].Next, &nodes[i + 1])
					nodes[i+1].Prev = &nodes[j]
					}
				}
			}


		reversDfs(dfs(&nodes[0]), seller, amount)

		for _, v := range nodes {
		fmt.Println(v)
		}
	return []int{}
}

func main()  {
	enroll := []string{"john", "mary", "edward", "sam", "emily", "jaimie", "tod", "young"}	
	referral := []string{"-", "-", "mary", "edward", "mary", "mary", "jaimie", "edward"}	
	seller := []string{"young", "john", "tod", "emily", "mary"}
	amount := []int{12, 4, 2, 5, 10}
	solution(enroll, referral, seller, amount)
}

func dfs(node *Node) []*Node {
	discoverd := make(map[string]int)
	stack := make([]*Node, 0)
	reseverNodes := make([]*Node, 0)
	reseverNodesPtr := &reseverNodes
	stack = append(stack, node)
	discoverd[node.Name] = 1
	stackPtr := &stack

 for; len(stack) != 0; {
	 nowNode := pop(stackPtr)

	 for _, v := range nowNode.Next {
		 if(discoverd[v.Name] == 0) {
			discoverd[v.Name] = 1 
			*reseverNodesPtr = append([]*Node{v}, reseverNodes...)
			*stackPtr = append(stack, v)
		 }
	 }
 }

return reseverNodes
}

func reversDfs(nodes []*Node,  seller []string, amount []int){
	fmt.Println(nodes)
	for i := 0; i < len(nodes); i++ {
		if nodes[i].MyMoney > 10 {
			tenPesent := nodes[i].MyMoney / 10
			nodes[i].MyMoney = nodes[i].MyMoney - tenPesent
			if nodes[i].Prev != nil {
				nodes[i].Prev.MyMoney = nodes[i].Prev.MyMoney + tenPesent
		}
	}
		for j := 0; j < len(seller); j++ {
			if  seller[j] == nodes[i].Name {
				money := amount[j] * 100
				if money > 10 && nodes[i].Prev != nil {
				 giveMoney :=  money / 10 
				 nodes[i].MyMoney = nodes[i].MyMoney + money - giveMoney
				 nodes[i].Prev.MyMoney = nodes[i].Prev.MyMoney + giveMoney
				} else {
					nodes[i].MyMoney = nodes[i].MyMoney + money 
				}
			}
		}
	 }
	}

func pop(nodes *[]*Node) *Node {
	popNode := (*nodes)[len(*nodes) - 1]
	*nodes = (*nodes)[:len(*nodes) - 1]

	return popNode
}