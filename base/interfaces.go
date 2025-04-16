package base

// ContainerInterface defines a common set of methods for container-like data structures.
//
// It provides basic introspection and management capabilities that are shared
// across stacks, queues, and other linear collections.
type ContainerInterface[T any] interface {
	// Len returns the number of elements currently in the container.
	Len() int

	// IsEmpty returns true if the container contains no elements.
	IsEmpty() bool

	// Clear removes all elements from the container. Capacity is retained.
	Clear()
}
