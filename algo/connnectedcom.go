package main

import "fmt"

func main() {
	fmt.Println("find connected component in a graph")
	var V int
	fmt.Scanln(&V)
	graph := make([][]bool, V)
	for i := 0; i < V; i++ {
		graph[i] = make([]bool, V)
	}

	var E int
	fmt.Scanln(&E)
	var src, dst int
	for i := 0; i < E; i++ {
		fmt.Scanln(&src, &dst)
		graph[src][dst] = true
	}
	visited := make([]bool, V)
	connComp := 0
	for i := 0; i < V; i++ {
		if !visited[i] {
			dfs(i, graph, visited)
			connComp++
		}
	}
	fmt.Println("connected components are ", connComp)
}

func dfs(s int, graph [][]bool, visited []bool) {
	visited[s] = true
	for idx := range graph[s] {
		if !graph[s][idx] {
			dfs(idx, graph, visited)
		}
	}
}
