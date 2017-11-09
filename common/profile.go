package common

import (
	"sync"
)

type (
	AtomUint struct {
		rw  sync.RWMutex
		val uint
	}

	CounterMap map[string]*AtomUint
	Profile    struct {
		rw      sync.RWMutex
		counter CounterMap
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

func (p *Profile) Get(name string) uint {
	p.rw.RLock()
	defer p.rw.RUnlock()
	cnt, ok := p.counter[name]
	if ok {
		return cnt.Get()
	}
	return 0
}

func NewProfile() *Profile {
	return &Profile{counter: make(CounterMap)}
}

func TheProfile() *Profile {
	once.Do(func() {
		instance = NewProfile()
	})
	return instance
}
