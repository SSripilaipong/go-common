package rslt

func JoinError[T error](x Of[T]) error {
	if x.IsErr() {
		return x.Error()
	}
	return x.Value()
}
