package fsa

type state int

func newStateSet(states ...state) stateSet {
	res := make(map[state]bool)
	for _, s := range states {
		res[s] = true
	}
	return res
}

type stateSet map[state]bool

func (set stateSet) has(s state) bool {
	return set[s]
}

func (set stateSet) add(s state) stateSet {
	set[s] = true
	return set
}

func (set stateSet) remove(s state) stateSet {
	set[s] = false
	return set
}

func (set stateSet) union(s2 stateSet) stateSet {
	res := make(map[state]bool)
	for s, ok := range set {
		if ok {
			res[s] = true
		}
	}
	for s, ok := range s2 {
		if ok {
			res[s] = true
		}
	}
	return res
}

func (set stateSet) copy() stateSet {
	res := make(map[state]bool)
	for s, ok := range set {
		if ok {
			res[s] = true
		}
	}
	return res
}

// util method for thompson construction
func (set stateSet) offset(i int) stateSet {
	res := make(map[state]bool)
	for s, ok := range set {
		if ok {
			res[s+state(i)] = true
		}
	}
	return res
}

func (set stateSet) size() int {
	i := 0
	for _, ok := range set {
		if ok {
			i++
		}
	}
	return i
}

func (set stateSet) elements() []state {
	res := make([]state, 0)
	for s, ok := range set {
		if ok {
			res = append(res, s)
		}
	}
	return res
}

func (set stateSet) intersection(s2 stateSet) stateSet {
	res := make(map[state]bool, 0)
	for s, ok := range set {
		if !ok {
			continue
		}
		ok = s2.has(s)
		if !ok {
			continue
		}
		res[s] = true
	}
	return res
}

func (set stateSet) equal(set2 stateSet) bool {
	for s, ok := range set {
		if !ok {
			continue
		}
		ok = set2.has(s)
		if !ok {
			return false
		}
	}

	for s, ok := range set2 {
		if !ok {
			continue
		}
		ok = set.has(s)
		if !ok {
			return false
		}
	}
	return true
}
