package goniq

import "sort"

func sortStringSlice(s []string) []string {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return s
}

func Add(slice []string, s string) []string {
	if len(slice) == 0 {
		return []string{s}
	}

	slice = sortStringSlice(slice)
	idx := sort.SearchStrings(slice, s)
	if idx == len(slice) || slice[idx] != s {
		return append(slice[:idx], append([]string{s}, slice[idx:]...)...)
	}
	return slice
}

func Remove(slice []string, s string) []string {
	if len(slice) == 0 {
		return slice
	}

	slice = sortStringSlice(slice)
	idx := sort.SearchStrings(slice, s)
	if idx < len(slice) && slice[idx] == s {
		if idx == len(slice)-1 {
			return slice[:idx]
		}
		copy(slice[idx:], slice[idx+1:])
		return slice[:len(slice)-1]
	}
	return slice
}
