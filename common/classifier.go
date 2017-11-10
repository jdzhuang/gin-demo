package common

import (
	"sort"
	"time"
)

type (
	Classifier interface {
		Declare(val interface{}, cls string)
		Deduce()
		Classify(val interface{}) string
		List() []string
	}

	DurationKnowledge  map[time.Duration]string
	DurationClassifier struct {
		knowledge  DurationKnowledge
		sorted_idx []time.Duration
	}
)

func NewDurationClassifier() *DurationClassifier {
	return &DurationClassifier{
		knowledge:  make(DurationKnowledge),
		sorted_idx: make([]time.Duration, 0),
	}
}

func (dc *DurationClassifier) Declare(val interface{}, cls string) {
	if v, ok := val.(time.Duration); ok {
		if _, exists := dc.knowledge[v]; !exists {
			dc.knowledge[v] = cls
			dc.sorted_idx = append(dc.sorted_idx, v)
		}
	}
}

func (dc *DurationClassifier) Deduce() {
	sort.SliceStable(dc.sorted_idx, func(i, j int) bool {
		return dc.sorted_idx[i] < dc.sorted_idx[j]
	})
}

func (dc *DurationClassifier) Size() int {
	return len(dc.sorted_idx)
}

func (dc *DurationClassifier) Hit(val interface{}) int {
	if v, ok := val.(time.Duration); ok {
		l := len(dc.sorted_idx)
		r := sort.Search(l, func(i int) bool {
			return dc.sorted_idx[i] >= v
		})
		//if v != dc.sorted_idx[i]
		//then r is the very index, that v should be insert with,
		return r
	}
	return -1
}

func (dc *DurationClassifier) Classify(val interface{}) string {
	idx := dc.Hit(val)
	l := len(dc.sorted_idx)
	switch {
	case idx >= 0 && idx < l:
		return dc.knowledge[dc.sorted_idx[idx]]
	case idx >= l:
		return dc.knowledge[dc.sorted_idx[l-1]] + "+"
	default:
		return "unknown class."
	}
}

func (dc *DurationClassifier) List() []string {
	slc := make([]string, len(dc.sorted_idx))
	for i, v := range dc.sorted_idx {
		slc[i] = dc.knowledge[v]
	}
	return slc
}
