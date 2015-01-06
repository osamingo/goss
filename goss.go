package goss

import "sort"

type (
	// Searcher is a interface for sort.Search
	Searcher interface {
		Target() int64
		Priority() string
	}

	// SortedSlice is a struct of sorted-slice
	SortedSlice struct {
		S    []Searcher
		DESC bool
	}
)

// Add append to slice
func (s *SortedSlice) Add(item Searcher) {

	if item == nil {
		return
	}

	i := sort.Search(len(s.S), func(i int) bool {
		if s.DESC {
			return s.S[i].Target() < item.Target() || (s.S[i].Target() == item.Target() && s.S[i].Priority() > item.Priority())
		}
		return s.S[i].Target() > item.Target() || (s.S[i].Target() == item.Target() && s.S[i].Priority() > item.Priority())
	})

	s.S = append(s.S, nil)
	copy(s.S[i+1:], s.S[i:])
	s.S[i] = item
}
