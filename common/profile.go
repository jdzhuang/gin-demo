package common

import (
	"fmt"
	"sync"
	"time"
)

type (
	AtomUint struct {
		rw  sync.RWMutex
		val uint
	}

	CounterMap map[string]*AtomUint
	Profile    struct {
		rw         sync.RWMutex
		counter    CounterMap
		classifier Classifier
	}
)

var (
	once     sync.Once
	instance *Profile
)

func (au *AtomUint) Inc(i uint) {
	au.rw.Lock()
	defer au.rw.Unlock()
	au.val = au.val + i
}

func (au *AtomUint) Get() uint {
	au.rw.RLock()
	defer au.rw.RUnlock()
	return au.val
}

func (p *Profile) Inc(name string, i uint) {
	p.rw.Lock()
	defer p.rw.Unlock()
	cnt, ok := p.counter[name]
	if !ok {
		cnt = new(AtomUint)
		p.counter[name] = cnt
	}
	cnt.Inc(i)
}

func (p *Profile) Since(start time.Time) {
	d := time.Since(start)
	cls := p.classifier.Classify(d)
	p.Inc(cls, 1)
}

func (p *Profile) Get(name string) uint {
	p.rw.RLock()
	defer p.rw.RUnlock()
	cnt, ok := p.counter[name]
	if ok {
		return cnt.Get()
	}
	return 0
}

func (p *Profile) Declare(val interface{}, cls string) {
	p.classifier.Declare(val, cls)
}

func (p *Profile) String() string {
	slc := p.classifier.List()
	slc = append(slc, slc[len(slc)-1]+"+")
	s := ""
	total := uint(0)
	for _, v := range slc {
		if cnt, ok := p.counter[v]; ok {
			total = total + cnt.Get()
			s = s + fmt.Sprintf("%-20s %d\n", v, cnt.Get())
		}
	}
	return s + fmt.Sprintf("%-20s %d\n", "total", total)
}

func NewProfile() *Profile {
	return &Profile{
		counter:    make(CounterMap),
		classifier: NewDurationClassifier(),
	}
}

func TheProfile() *Profile {
	once.Do(func() {
		instance = NewProfile()
	})
	return instance
}
