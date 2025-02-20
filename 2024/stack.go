// Node representa un nodo en la pila genérica.
type Node[T any] struct {
	value T
	next  *Node[T]
}

// Stack representa una pila genérica.
type Stack[T any] struct {
	top  *Node[T]
	size int
}

// Push agrega un elemento a la pila en O(1).
func (s *Stack[T]) Push(value T) {
	newNode := &Node[T]{value: value, next: s.top}
	s.top = newNode
	s.size++
}

// Pop elimina y devuelve el elemento superior de la pila en O(1).
// Si la pila está vacía, devuelve el valor por defecto del tipo y false.
func (s *Stack[T]) Pop() (T, bool) {
	if s.top == nil {
		var zero T // Valor por defecto del tipo genérico.
		return zero, false
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value, true
}

// Peek devuelve el elemento superior sin eliminarlo en O(1).
// Si la pila está vacía, devuelve el valor por defecto del tipo y false.
func (s *Stack[T]) Peek() (T, bool) {
	if s.top == nil {
		var zero T
		return zero, false
	}
	return s.top.value, true
}

// IsEmpty verifica si la pila está vacía.
func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

// Size devuelve el número de elementos en la pila.
func (s *Stack[T]) Size() int {
	return s.size
}

// str
func (s *Stack[T]) String() string {
	str := ""
	node := s.top
	for node != nil {
		str += fmt.Sprintf("%v,", node.value)
		node = node.next
		if node != nil {
			str += " "
		}
	}
	return str
}