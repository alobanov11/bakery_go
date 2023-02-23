package bake

import "time"

type Model struct {
	time time.Duration
	name string
}

func New(t time.Duration, n string) *Model {
	m := &Model{
		time: t,
		name: n,
	}
	return m
}

func (m *Model) CookingTime() time.Duration {
	return m.time
}

func (m *Model) Name() string {
	return m.name
}
