package utils

type Ease struct {
	From float64
	To   float64
}

func (e Ease) Get(t float64) float64 {
	return e.From + (e.To-e.From)*t
}

type EaseOut struct {
	From float64
	To   float64
}

func (e EaseOut) Get(t float64) float64 {
	return e.From + (e.To-e.From)*t*t
}
