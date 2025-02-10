package structures

import (
	"fmt"
)

// Grafo representado con lista de adyacencia
type Grafo struct {
	adj map[int][]int // Mapa de nodos y sus conexiones
}

// NewGrafo crea un nuevo grafo vacío
func NewGrafo() *Grafo {
	return &Grafo{adj: make(map[int][]int)}
}

// AgregarNodo añade un nodo al grafo
func (g *Grafo) AgregarNodo(nodo int) {
	if _, existe := g.adj[nodo]; !existe {
		g.adj[nodo] = []int{}
	}
}

// AgregarArista añade una conexión entre dos nodos
func (g *Grafo) AgregarArista(origen, destino int) {
	g.AgregarNodo(origen)
	g.AgregarNodo(destino)
	g.adj[origen] = append(g.adj[origen], destino)
}

// EliminarArista elimina una arista entre dos nodos
func (g *Grafo) EliminarArista(origen, destino int) {
	if _, existe := g.adj[origen]; existe {
		nuevaLista := []int{}
		for _, v := range g.adj[origen] {
			if v != destino {
				nuevaLista = append(nuevaLista, v)
			}
		}
		g.adj[origen] = nuevaLista
	}
}

// EliminarNodo elimina un nodo y todas sus conexiones
func (g *Grafo) EliminarNodo(nodo int) {
	delete(g.adj, nodo)
	for k := range g.adj {
		g.EliminarArista(k, nodo)
	}
}

// BFS - Recorrido en anchura (Breadth-First Search)
func (g *Grafo) BFS(inicio int) {
	visitados := make(map[int]bool)
	cola := []int{inicio}

	fmt.Print("BFS: ")
	for len(cola) > 0 {
		nodo := cola[0]
		cola = cola[1:]

		if visitados[nodo] {
			continue
		}

		fmt.Print(nodo, " ")
		visitados[nodo] = true

		for _, vecino := range g.adj[nodo] {
			if !visitados[vecino] {
				cola = append(cola, vecino)
			}
		}
	}
	fmt.Println()
}

// DFS - Recorrido en profundidad (Depth-First Search)
func (g *Grafo) DFS(inicio int) {
	visitados := make(map[int]bool)
	var dfsRecursivo func(nodo int)

	dfsRecursivo = func(nodo int) {
		if visitados[nodo] {
			return
		}
		fmt.Print(nodo, " ")
		visitados[nodo] = true
		for _, vecino := range g.adj[nodo] {
			dfsRecursivo(vecino)
		}
	}

	fmt.Print("DFS: ")
	dfsRecursivo(inicio)
	fmt.Println()
}

// Mostrar imprime el grafo
func (g *Grafo) Mostrar() {
	fmt.Println("Grafo:")
	for nodo, vecinos := range g.adj {
		fmt.Printf("%d -> %v\n", nodo, vecinos)
	}
}

func Entrypoint() {
	g := NewGrafo()

	// Agregar nodos y aristas
	g.AgregarArista(1, 2)
	g.AgregarArista(1, 3)
	g.AgregarArista(2, 4)
	g.AgregarArista(3, 5)
	g.AgregarArista(4, 5)
	g.AgregarArista(5, 1) // Crea un ciclo

	// Mostrar el grafo
	g.Mostrar()

	// Realizar BFS y DFS desde el nodo 1
	g.BFS(1)
	g.DFS(1)

	// Eliminar un nodo y mostrar el grafo actualizado
	g.EliminarNodo(3)
	g.Mostrar()
}
