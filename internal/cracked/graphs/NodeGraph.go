package graphs

import "fmt"

type void struct{}

var member void

type Graph struct {
	nodez map[string]map[string]void // map of sets
}

func (g *Graph) addDirectedEdge(src, dst string) {
	if existingSet, exists := g.nodez[src]; exists { // edge src -> dst DNE
		existingSet[dst] = member
	} else {
		dstSet := make(map[string]void)
		dstSet[dst] = member
		g.nodez[src] = dstSet
	}
}

func (g *Graph) removeDirectedEdge(src, dst string) {
	if existingSet, exists := g.nodez[src]; exists {
		delete(existingSet, dst)
	} // else - no-op
}

func (g *Graph) RouteExists(src, dst string) bool {
	var visited = make(map[string]void)
	visited[src] = member
	// Does src even exists
	if _, exists := g.nodez[src]; !exists {
		return false
	} else {
		if adjacent, exists := g.nodez[src]; exists {
			for s := range adjacent {
				visited[s] = member
				fmt.Println(s)
			}
		}

	}
	return false
}

func doesRouteExist(src, dst string, graph map[string]map[string]void) bool {
	return false
}

func contains(set map[string]void, tgt string) (exists bool) {
	_, exists = set[tgt]
	return
}

func DepthFirstSearch(g Graph, tgt string) bool {
	return false
}
