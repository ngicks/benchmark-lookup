package benchmarklookup

func LookUpSlice[T comparable](tab []T, target T) int {
	for idx, val := range tab {
		if val == target {
			return idx
		}
	}

	return -1
}
