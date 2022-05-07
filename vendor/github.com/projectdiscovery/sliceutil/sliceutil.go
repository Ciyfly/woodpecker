package sliceutil

// PruneEmptyStrings from the slice
func PruneEmptyStrings(v []string) []string {
	return PruneEqual(v, "")
}

// PruneEqual removes items from the slice equal to the specified value
func PruneEqual(v []string, equalTo string) (r []string) {
	for i := range v {
		if v[i] != equalTo {
			r = append(r, v[i])
		}
	}
	return
}

// Dedupe removes duplicates from a slice of strings preserving the order
func Dedupe(v []string) (r []string) {
	seen := make(map[string]struct{})
	for _, vv := range v {
		if _, ok := seen[vv]; !ok {
			seen[vv] = struct{}{}
			r = append(r, vv)
		}
	}
	return
}

func DedupeInt(v []int) (r []int) {
	seen := make(map[int]struct{})
	for _, vv := range v {
		if _, ok := seen[vv]; !ok {
			seen[vv] = struct{}{}
			r = append(r, vv)
		}
	}
	return
}
