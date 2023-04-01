package list

type ListElement[Key comparable, Value any] struct {
	keyHash uintptr
}
