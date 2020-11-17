package common

type Lib struct {
}

func (l *Lib) CompareStable(i, j int) bool {
	return i > j
}

func (l *Lib) Compare(i, j int) bool {
	return i >= j
}

func (l *Lib) Swap(i, j int) {
	i, j = j, i
}
