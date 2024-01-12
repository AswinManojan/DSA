package main

import "fmt"

type GraphNode struct {
	Vertex map[int][]int
}

func (g *GraphNode) Insert(vertex, edge int, isBi bool) {
	g.Vertex[vertex] = append(g.Vertex[vertex], edge)
	if isBi {
		g.Vertex[edge] = append(g.Vertex[edge], vertex)
	}
}

func (g *GraphNode) Display() {
	for vertex, edge := range g.Vertex {
		fmt.Println(vertex, edge)
	}
}

func (g *GraphNode) RemoveEdge(v1, v2 int, isBi bool) {
	arr := g.Vertex[v1]
	for i, v := range arr {
		if v == v2 {
			g.Vertex[v1] = append(arr[:i], arr[i+1:]...)
			break
		}
	}
	if len(g.Vertex[v1]) == 0 {
		delete(g.Vertex, v1)
	}
	if isBi {
		arr := g.Vertex[v2]
		for i, v := range arr {
			if v == v1 {
				g.Vertex[v2] = append(arr[:i], arr[i+1:]...)
				break
			}
		}
		if len(g.Vertex[v2]) == 0 {
			delete(g.Vertex, v2)
		}
	}
}

func (g *GraphNode) RemoveVertex(v int) {
	delete(g.Vertex, v)
	for key, arr := range g.Vertex {
		for i, e := range arr {
			if e == v {
				g.Vertex[key] = append(arr[:i], arr[i+1:]...)
			}
		}
	}
}
func (g *GraphNode) BFS(value int) {
	visited := make(map[int]bool)
	visited[value] = true
	arr := []int{value}
	for len(arr) > 0 {
		val := arr[0]
		fmt.Println(val)
		arr = arr[1:]
		for _, x := range g.Vertex[val] {
			if !visited[x] {
				visited[x] = true
				arr = append(arr, x)
			}
		}
	}
}

func (g *GraphNode) DFS(value int) {
	visited := make(map[int]bool)
	visited[value] = true
	stack := []int{value}
	g.DFSHelper(value, visited, stack)
}

func (g *GraphNode) DFSHelper(value int, visited map[int]bool, stack []int) {
	if len(stack) == 0 {
		return
	}
	val := stack[len(stack)-1]
	fmt.Println(val)
	stack = stack[:len(stack)-1]
	for _, v := range g.Vertex[val] {
		if !visited[v] {
			stack = append(stack, v)
			visited[v] = true
			g.DFSHelper(v, visited, stack)
		}
	}
}

func main() {
	g := GraphNode{Vertex: make(map[int][]int)}

	g.Insert(1, 2, false)
	g.Insert(2, 3, false)
	g.Insert(3, 4, true)
	g.Insert(2, 5, false)
	g.Insert(1, 5, false)
	g.Insert(6, 7, false)
	g.Insert(6, 2, false)
	g.Insert(3, 2, false)
	g.Insert(4, 5, true)

	g.RemoveEdge(4, 5, true)
	g.RemoveVertex(3)
	g.Display()
	g.DFS(2)
	g.BFS(2)
}
