package filter

type Interface interface {
	Len() int
	Index(int) string
	Remove(int)
}

type MatcherFunc func(string, string) bool

func Filter(source string, targets Interface, match MatcherFunc) {
	for i := 0; i < targets.Len(); {
		if match(source, targets.Index(i)) {
			i++
		} else {
			targets.Remove(i)
		}
	}
}
