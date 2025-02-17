package datastructures

var ValueArista int = 0

type Arista struct {
	Value int
}

type Graph struct {
	Arista map[string][]Arista
}

func New() *Graph {
	return &Graph{
		Arista: make(map[string][]Arista),
	}
}

func (g *Graph) AddNode(node string) {
	if _, ok := g.Arista[node]; !ok {
		ValueArista++
		g.Arista[node] = []Arista{
			{Value: ValueArista},
		}
	}
}

func (g *Graph) AddVertice(v1, v2 string) {
	g.AddNode(v1)
	g.AddNode(v2)

	g.Arista[v1] = append(g.Arista[v1], v2)
	g.Arista[v2] = append(g.Arista[v2], v1)
}

func (g *Graph) DeleteNode() {}
