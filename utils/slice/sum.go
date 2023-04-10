package slice

func SumFunc[T any, M any](slice []T, sum func(M, T) M) M {
	var result M
	for _, v := range slice {
		result = sum(result, v)
	}
	return result
}
