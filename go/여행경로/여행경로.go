package main

import "fmt"

type Node struct {
	Name string
	to   []*Node // 2개 이상일 수 있음.
}

func solution(tickets [][]string) []string {
	nodeMap := make(map[string]*Node)
	checkMap := make(map[string]int)

	for _, v := range tickets {
		if checkMap[v[0]] == 0 {
			fromNode := Node{Name: v[0]}
			checkMap[fromNode.Name] = 1
			nodeMap[fromNode.Name] = &fromNode
		}
		if checkMap[v[1]] == 0 {
			toNode := Node{Name: v[1]}
			checkMap[toNode.Name] = 1
			nodeMap[toNode.Name] = &toNode
		}
	}

	for _, v := range tickets {
		from, to := v[0], v[1]
		node := nodeMap[from]
		node.to = append(node.to, nodeMap[to])
	}

	usedTickets := make(map[string]int)

	q := []*Node{nodeMap["ICN"]}
	qPtr := &q
	visitedAirport := make([]string, 0)
	visitedAirportPtr := &visitedAirport

	for i := 0; i < len(tickets); i++ {

		initNode := Node{Name: "ZZZ"}		
		node := q[0]
		(*qPtr)[0] = &initNode

		for _, v := range node.to {
			if int(q[0].Name[0])+int(q[0].Name[1])+int(q[0].Name[2]) > int(v.Name[0])+int(v.Name[1])+int(v.Name[2]) && usedTickets[node.Name+v.Name] != 1 {
				(*qPtr)[0] = v
			}
		}

		if q[0].Name == "ZZZ"{
			continue
		}

		usedTickets[node.Name+q[0].Name] = 1

		if len(visitedAirport) != 0 {
			(*visitedAirportPtr)[len(visitedAirport)-1] = node.Name
			*visitedAirportPtr = append(visitedAirport, q[0].Name)
		} else {
			*visitedAirportPtr = append(visitedAirport, node.Name, q[0].Name)
		}
	}

	return visitedAirport
}

func main() {
	test := [][]string{{"ICN", "SFO"}, {"SFO", "ICN"}, {"ICN", "SFO"}, {"SFO", "QRE"}}

	fmt.Println(solution(test))
}