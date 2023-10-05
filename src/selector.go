package src

import (
	"fmt"
	"math/rand"
	"sync"
)

type Selector interface {
	Select(words []string) string
}

type RandomSelector struct{}

func (r *RandomSelector) Select(words []string) string {
	idx := rand.Intn(len(words))
	return words[idx]
}

type ISelectorRegistry interface {
	Get(selector string) (Selector, error)
	Register(name string, selector Selector) error
}

type LocalSelectorRegistry struct {
	sync.Mutex
	registry map[string]Selector
}

func (sr *LocalSelectorRegistry) Get(name string) (Selector, error) {
	sr.Lock()
	defer sr.Unlock()
	sel, ok := sr.registry[name]
	if !ok {
		return nil, fmt.Errorf("name %v not found", name)
	}
	return sel, nil
}

func (sr *LocalSelectorRegistry) Register(name string, selector Selector) error {
	sr.Lock()
	defer sr.Unlock()
	if _, ok := sr.registry[name]; ok {
		return fmt.Errorf("selector with name %v already exists", name)
	}
	sr.registry[name] = selector
	return nil
}
