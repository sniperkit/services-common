package sets

type StringSet struct {
	set map[string]bool
}

func NewStringSet() *StringSet {
	return &StringSet{make(map[string]bool)}
}

func (set *StringSet) Add(s string) bool {
	_, found := set.set[s]
	set.set[s] = true
	return !found
}

func (set *StringSet) Remove(s string) {
	delete(set.set, s)
}

func (set *StringSet) Contains(s string) bool {
	_, found := set.set[s]
	return found
}

func (set *StringSet) Size() int {
	return len(set.set)
}

func (set *StringSet) FromSlice(slice []string) {
	for s := range slice {
		set.Add(slice[s])
	}
}

func (set *StringSet) Intersection(slice1 *StringSet, slice2 *StringSet) *StringSet {
	var output = NewStringSet()

	for s := range slice1.set {
		if slice2.Contains(s) {
			output.Add(s)
		}
	}

	return output
}

func (set *StringSet) Difference(slice1 *StringSet, slice2 *StringSet) *StringSet {
	var output = NewStringSet()

	for s := range slice1.set {
		if !slice2.Contains(s) {
			output.Add(s)
		}
	}

	return output
}
