package grains

func Square(in int) (uint64, error) {
	var sum uint64
	if in == 1 {
		return uint64(in), nil
	}
	val, err := Square(in-1)
	if err != nil {
		return 1, err
	}
	sum += val
	return sum, nil
}

func Total() uint64 {
	return 1
}