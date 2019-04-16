package fsa

func newIntegerSet() *integerSet {
	var res integerSet = make(map[int]bool)
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

func (i *integerSet) equals(i2 *integerSet) bool {
	flag := true
	for _, e := range i.entries() {
		flag = flag && i2.has(e)
		if !flag {
			return false
		}
	}
	for _, e := range i2.entries() {
		flag = flag && i.has(e)
		if !flag {
			return false
		}
	}
	return flag
}
