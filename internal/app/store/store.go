package store

import (
	"bakery/internal/app/bake"
	"bakery/internal/app/factory"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

type Store interface {
	Prepare(t bake.Type) (*bake.Model, error)
	Cook(b *bake.Model)
	Types() []bake.Type
	Result() string
}

type store struct {
	factory factory.Factory
	types   []bake.Type
	cap     int
	cur     *sync.Map
	res     *sync.Map
}

func New(file string, cap int, factory factory.Factory) (Store, error) {
	jsonFile, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	var types []bake.Type
	err = json.Unmarshal(jsonFile, &types)

	if err != nil {
		return nil, err
	}

	s := &store{
		factory: factory,
		types:   types,
		cap:     cap,
		cur:     &sync.Map{},
		res:     &sync.Map{},
	}

	return s, nil
}

func (s *store) Prepare(t bake.Type) (*bake.Model, error) {
	v, ok := s.cur.Load(t.Name)

	if !ok {
		return s.factory.New(t), nil
	}

	count := v.(int)

	if ok && count == s.cap {
		return nil, errors.New("the type is reached max number to bake")
	}

	return s.factory.New(t), nil
}

func (s *store) Cook(b *bake.Model) {
	s.retainBake(b.Name())
	time.Sleep(b.CookingTime())
	s.releaseBake(b.Name())
}

func (s *store) Types() []bake.Type {
	return s.types
}

func (s *store) Result() string {
	r := ""

	s.res.Range(func(key, value any) bool {
		v := fmt.Sprintf("%v(%v)", key, value)
		if r == "" {
			r = v
		} else {
			r = r + ", " + v
		}
		return true
	})

	return r
}

func (s *store) retainBake(b string) {
	v, ok := s.cur.Load(b)
	count := 0
	if ok {
		count = v.(int)
	}
	s.cur.Store(b, count+1)
}

func (s *store) releaseBake(b string) {
	v, ok := s.cur.Load(b)

	count := 0
	if ok {
		count = v.(int)
	}

	if count == 1 {
		s.cur.Delete(b)
	} else {
		s.cur.Store(b, count-1)
	}

	v, ok = s.res.Load(b)

	count = 0
	if ok {
		count = v.(int)
	}

	s.res.Store(b, count+1)
}
