package models

type Number interface {
	int | float64
}

type Calculator[T Number] struct {
	A T
	B T
}

func (c Calculator[T]) Adding() T {
	return c.A + c.B
}

func (c Calculator[T]) Substracting() T {
	return c.A - c.B
}

func (c Calculator[T]) Multiply() T {
	return c.A * c.B
}

func (c Calculator[T]) Divide() T {
	return c.A / c.B
}
