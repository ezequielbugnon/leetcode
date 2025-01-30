package twopointers

/*
   1. Suma Objetivo en Array Ordenado
   Dado un array de números enteros ordenado en orden ascendente y un número objetivo target, encuentra dos índices (1-based) de los números que suman target.

   Ejemplo:
   Input: nums = [2, 7, 11, 15], target = 9
   Output: [1, 2]
   Explicación: nums[1] + nums[2] = 2 + 7 = 9
   Restricciones:

   Debe usarse un enfoque de two-pointers.
   Solo hay una solución única.
   No puedes usar la misma posición dos veces.
*/

func SumArrayHashMap(nums []int, target int) []int {
	visited := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, ok := visited[complement]; ok {
			return []int{index, i}
		}

		visited[num] = i
	}
	return nil
}

func SumArrayTwoPointers(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1} // Devuelve en formato 1-based
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil // No hay solución
}

/*
   2. Cuadrados de un Array Ordenado
   Dado un array de enteros nums ordenado de forma no decreciente (puede contener negativos), retorna un nuevo array con los cuadrados de cada número ordenados también en orden no decreciente.

   Ejemplo:
   Input: nums = [-4, -1, 0, 3, 10]
   Output: [0, 1, 9, 16, 100]
   Restricciones:

   Usa two-pointers para resolver el problema en O(n).
   No puedes usar .sort() en la solución.

*/

func OrderCuadraticArray(nums []int) []int {
	n := len(nums)
	left, right := nums[0], n-1
	result := make([]int, n)
	index := n - 1

	for left <= right {
		leftSq := nums[left] * nums[left]
		rightSq := nums[right] * nums[right]

		if leftSq > rightSq {
			result[index] = leftSq
			left++
		} else {
			result[index] = rightSq
			right--
		}
		index--
	}

	return result
}

/*
       3. Eliminar Duplicados en un Array Ordenado (Medio)
   📌 Descripción:
   Dado un array ordenado nums, modifica el array in-place para que cada número aparezca solo una vez y devuelve la nueva longitud del array.

   📌 Ejemplo:

   plaintext
   Copiar
   Editar
   Entrada: nums = [0,0,1,1,1,2,2,3,3,4]
   Salida: 5  // nums = [0,1,2,3,4,_,_,_,_,_]
   📌 Pistas:

   Usa un puntero lento para sobrescribir valores repetidos.
   Un puntero rápido recorre el array y copia valores únicos en el puntero lento.

*/

/*
   4. Contenedor con Más Agua (Difícil)
   📌 Descripción:
   Dado un array height donde cada índice representa la altura de un contenedor, encuentra el área máxima de agua que se puede almacenar.

   📌 Ejemplo:

   plaintext
   Copiar
   Editar
   Entrada: height = [1,8,6,2,5,4,8,3,7]
   Salida: 49
   📌 Pistas:

   Usa dos punteros: uno al inicio y otro al final.
   Calcula el área y mueve el puntero que tiene menor altura para intentar maximizar el área.

*/
