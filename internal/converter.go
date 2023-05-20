package internal

func M(value float64) float64 {
	return value * 1000000
}

func G(value float64) float64 {
	return value * 1000000000
}

func T(value float64) float64 {
	return value * 1000000000000
}
