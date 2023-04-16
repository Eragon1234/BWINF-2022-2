package pkg

// Permutation is a struct that can be used to generate all permutations of a slice.
// The permutations don't include the original slice.
type Permutation struct {
	original    []int
	permutation []int
}

func NewPermutation(s []int) Permutation {
	return Permutation{
		original:    s,
		permutation: make([]int, len(s)),
	}
}

// Next returns the next permutation or nil if there are no more permutations.
// The permutations don't include the original.
func (p *Permutation) Next() []int {
	nextPerm(p.permutation)
	ok := p.permutation[0] < len(p.permutation)
	if !ok {
		return nil
	}
	return getPerm(p.original, p.permutation)
}

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}
