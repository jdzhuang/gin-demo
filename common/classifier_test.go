package common

import (
	//	"sort"
	"testing"
	"time"
)

const (
	CLS20    = "20MS"
	CLS50    = "50MS"
	CLS100   = "100MS"
	CLSLARGE = "100MS+"
)

func PrepareDurationClassifier() *DurationClassifier {
	duration_cls := []struct {
		val time.Duration
		cls string
	}{
		{20 * time.Millisecond, CLS20},
		{50 * time.Millisecond, CLS50},
		{100 * time.Millisecond, CLS100},
	}
	dc := NewDurationClassifier()
	for _, v := range duration_cls {
		dc.Declare(v.val, v.cls)
	}
	dc.Deduce()
	return dc
}

func TestDurationClassifier_Hit(t *testing.T) {
	input := []struct {
		val     time.Duration
		exp_idx int
		desc    string
	}{
		{10 * time.Millisecond, 0, "10 ms"},
		{20 * time.Millisecond, 0, "20 ms"},
		{30 * time.Millisecond, 1, "30 ms"},
		{50 * time.Millisecond, 1, "50 ms"},
		{60 * time.Millisecond, 2, "60 ms"},
		{120 * time.Millisecond, 3, "120 ms"},
	}

	dc := PrepareDurationClassifier()
	for _, v := range input {
		idx := dc.Hit(v.val)
		if idx != v.exp_idx {
			t.Errorf("index['%s'] is not '%d' (%d)", v.desc, v.exp_idx, idx)
		}
	}
}

func TestDurationClassifier_Classify(t *testing.T) {
	input := []struct {
		val     time.Duration
		exp_cls string
		desc    string
	}{
		{10 * time.Millisecond, CLS20, "10 ms"},
		{20 * time.Millisecond, CLS20, "20 ms"},
		{30 * time.Millisecond, CLS50, "30 ms"},
		{50 * time.Millisecond, CLS50, "50 ms"},
		{60 * time.Millisecond, CLS100, "60 ms"},
		{120 * time.Millisecond, CLSLARGE, "120 ms"},
	}
	//
	dc := PrepareDurationClassifier()
	for _, v := range input {
		cls := dc.Classify(v.val)
		if cls != v.exp_cls {
			t.Errorf("'%s' is not classified as '%s' (%s)", v.desc, v.exp_cls, cls)
		}
	}

}

func TestDurationClassifier_Size(t *testing.T) {
	dc := PrepareDurationClassifier()
	size := dc.Size()
	t.Logf("size:%d.", size)
	dc.Declare(111*time.Second, "111 sec")
	if size+1 != dc.Size() {
		t.Errorf("size does not increase.")
		t.FailNow()
	}
	size = dc.Size()
	t.Logf("size:%d.", size)
	dc.Declare(111*time.Second, "111 sec again")
	if size < dc.Size() {
		t.Errorf("with same declaration, size should not increase.")
		t.FailNow()
	}
}

func TestNothing(t *testing.T) {
}
