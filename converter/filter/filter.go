package filter

import (
	"path/filepath"
	"strconv"
	"strings"
)

// Filter defines filter rules
type Filter interface {
	Apply(filenames []string) []string
}

func NewAlgorithmsFilter() Filter {
	return &AlgorithmsFilter{}
}

// AlgorithmsFilter defines algorithm
type AlgorithmsFilter struct{}

// Apply
func (f *AlgorithmsFilter) Apply(filenames []string) (filtered []string) {
	for _, filename := range filenames {
		subs := strings.Split(filepath.Base(filename), "_")
		_, err := strconv.ParseInt(subs[0], 10, 64)
		if err == nil {
			filtered = append(filtered, filename)
		}
	}
	return
}
