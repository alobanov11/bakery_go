package factory

import (
	"bakery/internal/app/bake"
	"time"
)

type Factory interface {
	New(t bake.Type) *bake.Model
}

type factory struct{}

func New() Factory {
	f := &factory{}
	return f
}

func (f *factory) New(t bake.Type) *bake.Model {
	return bake.New(time.Duration(t.Time)*time.Second, t.Name)
}
