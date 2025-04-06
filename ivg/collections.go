package ivg

import (
	"maps"
	"slices"
	"sync"
)

type IconResource struct {
	IconName string
	IconData []byte
}

func (ir *IconResource) Name() string {
	return ir.IconName
}
func (ir *IconResource) Content() []byte {
	return ir.IconData
}

type Collection map[string]*IconResource

type collections struct {
	mux sync.RWMutex
	c   map[string]Collection
}

func (cs *collections) Lookup(collection, icon string) *IconResource {
	cs.mux.RLock()
	defer cs.mux.RUnlock()
	coll, ok := cs.c[collection]
	if !ok {
		return nil
	}
	res, ok := coll[icon]
	if !ok {
		return nil
	}
	return res
}

func (cs *collections) Registered(name string, c Collection) Collection {
	cs.mux.Lock()
	defer cs.mux.Unlock()
	if cs.c == nil {
		cs.c = make(map[string]Collection)
	}
	cs.c[name] = c
	return c
}

func (cs *collections) Keys() map[string][]string {
	cs.mux.RLock()
	defer cs.mux.RUnlock()
	ret := make(map[string][]string)
	for cName := range maps.Keys(cs.c) {
		ret[cName] = slices.Collect(maps.Keys(cs.c[cName]))
	}
	return ret
}

var Collections = new(collections)
