package optional

type Of[T any] struct {
	x *T
}

func Value[T any](x T) Of[T] {
	return Of[T]{x: &x}
}

func ValueOf[T any](x Of[T]) T {
	return x.Value()
}

func Empty[T any]() Of[T] {
	return Of[T]{x: nil}
}

func (o Of[T]) Value() (x T) {
	if o.x == nil {
		return x
	}
	return *o.x
}

func (o Of[T]) IsEmpty() bool {
	return o.x == nil
}

func (o Of[T]) IsNotEmpty() bool {
	return !o.IsEmpty()
}

func (o Of[T]) Return() (x T, ok bool) {
	if o.x == nil {
		return x, false
	}
	return *o.x, true
}

func New[T any](x T, exists bool) Of[T] {
	if exists {
		return Value(x)
	}
	return Empty[T]()
}
