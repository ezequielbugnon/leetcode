package datastructures

type Edge struct {
	Value   int
	Destiny string
}

type Graph struct {
	Edge map[string][]Edge
}

func New() *Graph {
	return &Graph{
		Edge: make(map[string][]Edge),
	}
}

func (g *Graph) AddNode(node string) {
	if _, ok := g.Edge[node]; !ok {
		g.Edge[node] = []Edge{}
	}
}

func (g *Graph) AddVertice(v1, v2 string) {
	g.AddNode(v1)
	g.AddNode(v2)

	g.Edge[v1] = append(g.Edge[v1], Edge{
		Destiny: v2,
		Value:   1,
	})
	g.Edge[v2] = append(g.Edge[v2], Edge{
		Destiny: v1,
	})
}

func (g *Graph) DeleteNode(node string) {
	delete(g.Edge, node)

	for k := range g.Edge {
		g.Edge[k] = g.DeleteEdges(k, node)
	}
}

func (g *Graph) DeleteEdges(node string, DeleteEdge string) []Edge {
	var newEdges []Edge
	for _, adj := range g.Edge[node] {
		if DeleteEdge != adj.Destiny {
			newEdges = append(newEdges, adj)
		}
	}

	return newEdges
}
