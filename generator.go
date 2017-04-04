package passgenerator

import (
	"math/rand"
	"unsafe"

	"github.com/DivineRapier/passgenerator/storage"
)

type Generator interface {
	storage.Storage
	Sum() Generator
	Dump() string
}

type PassGenerator struct {
	container [][]byte
	length    int
	password  []byte
	storage.Storage
}

func NewPassGenerator() Generator {
	var generator PassGenerator
	for _, val := range *kinds {
		switch {
		case val >= '0' && val <= '9':
			generator.container = append(generator.container, digit)
		case val >= 'a' && val <= 'z':
			generator.container = append(generator.container, lower)
		case val >= 'A' && val <= 'Z':
			generator.container = append(generator.container, upper)
		default:
			generator.container = append(generator.container, symbol)
		}
	}
	generator.length = *length
	generator.password = make([]byte, *length, *length)
	generator.Storage = storage.NewFileStorage(*key)
	return &generator
}

func (g *PassGenerator) Sum() Generator {
	for i := 0; i < g.length; i++ {
		index := rand.Int() % len(g.container)
		g.password[i] = g.container[index][rand.Int()%len(g.container[index])]
	}
	return g
}

func (g *PassGenerator) Dump() string {
	return *(*string)(unsafe.Pointer(&g.password))
}
