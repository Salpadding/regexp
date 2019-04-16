package nfa

type integerSet map[int]bool

type transitions []transition

type states []int

func (s *states) offset(offset int) *states {
	var res states = make([]int, len(*s))
	for idx, i := range *s {
		res[idx] = i + offset
	}
	return &res
}

func (s *states) len() int {
	return len(*s)
}

func (s *states) concat(s2 *states) *states {
	var res states
	res = append(*s, *s2...)
	return &res
}

func (s *states) slice(start, end int) *states {
	var res states
	res = (*s)[start:end]
	return &res
}

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

func (i *integerSet) offset(offset int) *integerSet {
	var res integerSet = make(map[int]bool, len(*i))
	for k, v := range *i {
		if v {
			res.add(k + offset)
		}
	}
	return &res
}

func (i *integerSet) len() int {
	var res int
	for _, v := range *i {
		if v {
			res++
		}
	}
	return res
}

type transitionSet map[int]*transitions

func (s *transitionSet) len() int {
	res := 0
	for k, v := range *s {
		if k != 0 && v != nil {
			res++
		}
	}
	return res
}

// WARNING: not purify
func (s *transitionSet) remove(val int) *transitions {
	removed := (*s)[val]
	(*s)[val] = nil
	return removed
}

func (s *transitionSet) has(val int) bool {
	t, ok := (*s)[val]
	return t != nil && ok
}

func (s *transitionSet) get(val int) *transitions {
	return (*s)[val]
}

func (s *transitionSet) add(key int, ts *transitions) *transitionSet {
	if (*s)[key] == nil {
		(*s)[key] = ts
		return s
	}
	oldts := *(*s)[key]
	newts := *ts
	newts = append(oldts, newts...)
	(*s)[key] = &newts
	return s
}

func (s *transitionSet) union(s2 *transitionSet) *transitionSet {
	var res transitionSet = make(map[int]*transitions, len(*s)+len(*s2))
	for k, v := range *s {
		if v == nil {
			continue
		}
		res.add(k, v)
	}
	for k, v := range *s2 {
		if v == nil {
			continue
		}
		res.add(k, v)
	}
	return &res
}

func (s *transitionSet) offset(offset int) *transitionSet {
	var res transitionSet = make(map[int]*transitions, len(*s))
	for k, v := range *s {
		if v != nil {
			res.add(k+offset, v.offset(offset))
		}
	}
	return &res
}

func (s *transitionSet) copy() *transitionSet {
	return s.offset(0)
}

func (t *transitions) offset(offset int) *transitions {
	var res transitions = make([]transition, len(*t))
	for idx, t := range *t {
		copied := t.copy()
		copied.setTo(t.to() + offset)
		res[idx] = copied
	}
	return &res
}
