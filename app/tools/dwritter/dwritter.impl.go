package dwritter

import (
	"errors"
	"log"
	"sync"

	"github.com/freekup/product-scrapper/app/entity"
)

type DataWritter struct {
	handlers map[string]DWHandlerInterface
}

func NewDataWritter() DataWritter {
	return DataWritter{
		handlers: make(map[string]DWHandlerInterface),
	}
}

func (dw *DataWritter) RegisterWritter(name string, handler DWHandlerInterface) error {
	if _, ok := dw.handlers[name]; ok {
		return errors.New("handler already registered")
	}

	dw.handlers[name] = handler
	return nil
}

func (dw *DataWritter) Write(loc []string, products []entity.Product) error {
	var wg sync.WaitGroup

	for _, l := range loc {
		if _, isExist := dw.handlers[l]; !isExist {
			log.Fatal("type", l, "is not supported")
		}

		wg.Add(1)

		go func(action string) {
			defer wg.Done()

			err := dw.handlers[action].Store(products)
			if err != nil {
				log.Fatal(err)
			}
		}(l)
	}

	wg.Wait()
	return nil
}
