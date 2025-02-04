package structures

import "fmt"

// Nodo de árbol binario
type NodeThree struct {
	value int
	left  *NodeThree
	right *NodeThree
}

// Método para insertar un nuevo nodo en el árbol
func (n *NodeThree) Insert(value int) {
	if value < n.value {
		if n.left == nil {
			n.left = &NodeThree{value: value}
		} else {
			n.left.Insert(value)
		}
	} else if value > n.value {
		if n.right == nil {
			n.right = &NodeThree{value: value}
		} else {
			n.right.Insert(value)
		}
	}
}

// Método para buscar un valor en el árbol
func (n *NodeThree) Search(value int) bool {
	if n == nil {
		return false
	}
	if value < n.value {
		return n.left.Search(value)
	} else if value > n.value {
		return n.right.Search(value)
	}
	return true
}

// Método para encontrar el nodo mínimo (más a la izquierda)
func (n *NodeThree) Min() *NodeThree {
	if n.left != nil {
		return n.left.Min()
	}
	return n
}

// Método para eliminar un nodo del árbol
func (n *NodeThree) Delete(value int) *NodeThree {
	if n == nil {
		return nil
	}

	// Si el valor es menor que el nodo actual, ir a la izquierda
	if value < n.value {
		n.left = n.left.Delete(value)
	} else if value > n.value { // Si es mayor, ir a la derecha
		n.right = n.right.Delete(value)
	} else { // Si el valor es igual al nodo actual, eliminar el nodo
		if n.left == nil && n.right == nil {
			return nil
		} else if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}
		// Si tiene dos hijos, reemplazar con el mínimo de la subárbol derecho
		minNode := n.right.Min()
		n.value = minNode.value
		n.right = n.right.Delete(minNode.value)
	}
	return n
}

// Recorrido en orden (in-order traversal)
func (n *NodeThree) InOrder() {
	if n != nil {
		n.left.InOrder()
		fmt.Print(n.value, " ")
		n.right.InOrder()
	}
}
