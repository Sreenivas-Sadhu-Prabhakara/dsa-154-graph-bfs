package main

import "fmt"

// BFS returns the breadth-first visit order over an undirected graph with n
// vertices (0-based), covering every component in vertex order.
func BFS(n int, edges [][2]int) []int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	visited := make([]bool, n)
	out := make([]int, 0, n)
	for s := 0; s < n; s++ {
		if visited[s] {
			continue
		}
		visited[s] = true // mark on enqueue, not on dequeue
		queue := []int{s}
		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			out = append(out, v)
			for _, w := range adj[v] {
				if !visited[w] {
					visited[w] = true
					queue = append(queue, w)
				}
			}
		}
	}
	return out
}

func main() {
	fmt.Println(BFS(5, [][2]int{{0, 1}, {0, 2}, {1, 3}}))
}
