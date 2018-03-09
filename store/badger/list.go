package badger

// StreamList is the list of streams on the server
type StreamList []string

// Contains returns whether the given name is in the list
func (l StreamList) Contains(name string) bool {
	for _, n := range l {
		if n == name {
			return true
		}
	}
	return false
}

// Remove removes the provided stream from the list
func (l StreamList) Remove(name string) (nl StreamList) {
	var index int
	var found = false

	for i, n := range l {
		if n == name {
			index = i
			found = true
			break
		}
	}
	if !found {
		return l
	}

	nl = append(l[:index], l[index+1:]...)
	return nl
}

// Add adds the provided stream to the list
func (l StreamList) Add(name string) (nl StreamList) {
	nl = append(l, name)
	return nl
}
