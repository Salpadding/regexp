package fsa

var epsilon rune = 0

type transitionFunction map[int]map[rune]int

type TransitionFunction interface {
	States() []int
	EpsilonClosure(int) *integerSet
	Transition(int, rune) int
	TransitionStates([]int, rune) []int
}

func (t *transitionFunction) States() []int {
	var res []int
	for k, _ := range *t {
		res = append(res, k)
	}
	return res
}

func (t *transitionFunction) Transition(in int, char rune) int {
	m, ok := (*t)[in]
	if !ok {
		return 0
	}
	return m[char]
}

func (t *transitionFunction) TransitionStates(in []int, char rune) []int {
	res := make([]int, len(in))
	for idx, i := range in {
		res[idx] = t.Transition(i, char)
	}
	return res
}

func (t *transitionFunction) EpsilonClosure(in int) *integerSet {
	var res integerSet = make(map[int]bool)
	res.add(in)
	m, ok := (*t)[in]
	if !ok {
		return &res
	}
	for start, ok := m[epsilon]; ok; start, ok = m[epsilon] {
		if res.has(start) {
			break
		}
		res.add(start)
	}
	return &res
}

type integerSet map[int]bool

func (i *integerSet) union(i2 *integerSet) *integerSet {
	var res integerSet = make(map[int]bool, len(*i)+len(*i2))
	for k, v := range *i {
		if v {
			res.add(k)
		}
	}
	for k, v := range *i2 {
		if v {
			res.add(k)
		}
	}
	return &res
}

func (i *integerSet) has(val int) bool {
	has, ok := (*i)[val]
	return has && ok
}

func (i *integerSet) add(val int) {
	(*i)[val] = true
}

func (i *integerSet) remove(val int) {
	(*i)[val] = false
}

func (i *integerSet) entries() []int {
	var res []int
	for k, v := range *i {
		if v {
			res = append(res, k)
		}
	}
	return res
}

func (i *integerSet) card() int {
	return len(i.entries())
}
