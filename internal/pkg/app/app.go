package app

import (
	"bakery/internal/app/bake"
	"bakery/internal/app/factory"
	"bakery/internal/app/store"
	"fmt"
	"log"
	"sync"
	"time"
)

type App struct {
	time int
	cap  int
	file string
}

func New(time int, cap int, file string) (*App, error) {
	a := &App{
		time: time,
		cap:  cap,
		file: file,
	}
	return a, nil
}

func (a *App) Start() {
	start := time.Now()
	defer fmt.Printf("Elapsed: %s\n", time.Since(start))

	f := factory.New()
	s, err := store.New(a.file, a.cap, f)

	if err != nil {
		log.Panic(err)
	}

	end := start.Add((time.Second * time.Duration(a.time)))
	wg := &sync.WaitGroup{}

	for _, t := range s.Types() {
		for i := 0; i < a.cap; i++ {
			wg.Add(1)

			go func(t bake.Type) {
				defer wg.Done()

				fmt.Printf("Start %s\n", t.Name)

				for {
					if time.Now().After(end) {
						break
					}

					b, err := s.Prepare(t)

					if err != nil {
						continue
					}

					s.Cook(b)

					fmt.Printf("Done %s\n", b.Name())
				}
			}(t)
		}
	}

	wg.Wait()
	fmt.Printf("Prepared: %s\n", s.Result())
}
