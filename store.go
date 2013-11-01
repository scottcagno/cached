// *
// * Copyright 2013, Scott Cagno, All rights reserved
// * BSD Licensed - sites.google.com/site/bsdc3license
// *
// * cached :: store.go :: database store
// *

package cached

import (
	"sync"
	"time"
)

// database store
type Store struct {
	db map[string]interface{}
	gc map[string]int64
	sync.Mutex
}

// InitStore returns a pointer to a new instance of store
func InitStore() *Store {
	st := &Store{
		db: make(map[string]interface{}),
		gc: make(map[string]int64),
	}
	go st.rungc()
	return st
}

// rungc runs the garbage collector for timed keys
func (self *Store) rungc() {
	if len(self.gc) > 0 {
		self.Lock()
		for k, ttl := range self.gc {
			if ttl <= time.Now().Unix() {
				delete(self.db, k)
				delete(self.gc, k)
			} else {
				break
			}
		}
		self.Unlock()
	}
	time.AfterFunc(time.Duration(1)*time.Second, func() { self.rungc() })
}

// add
func (self *Store) add(r *Request) interface{} {
	return "hit add"
}

// set
func (self *Store) set(r *Request) interface{} {
	return "hit set"
}

// get
func (self *Store) get(r *Request) interface{} {
	return "hit get"
}

// del
func (self *Store) del(r *Request) interface{} {
	return "hit del"
}

// exp
func (self *Store) exp(r *Request) interface{} {
	return "hit exp"
}

// ttl
func (self *Store) ttl(r *Request) interface{} {
	return "hit ttl"
}
