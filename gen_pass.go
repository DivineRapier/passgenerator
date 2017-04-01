package passgenerator

import (
	"flag"
	"math/rand"
	"time"
	"unsafe"
)

type passGenerator struct {
	container [][]byte
	length    int
	password  []byte
}

var (
	digit = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	lower = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	upper = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
)

var (
	set    = flag.String("s", "1aA", "kinds of letter make up password")
	length = flag.Int("l", 16, "length of password")
)

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
}

func initGenerator() *passGenerator {
	var passGen passGenerator
	for _, val := range *set {
		switch {
		case val >= '0' && val <= '9':
			passGen.container = append(passGen.container, digit)
		case val >= 'a' && val <= 'z':
			passGen.container = append(passGen.container, lower)
		case val >= 'A' && val <= 'Z':
			passGen.container = append(passGen.container, upper)
		}
	}
	passGen.length = *length
	passGen.password = make([]byte, *length, *length)
	return &passGen
}

func (g *passGenerator) sum() *passGenerator {
	for i := 0; i < g.length; i++ {
		index := rand.Int() % len(g.container)
		g.password[i] = g.container[index][rand.Int()%len(g.container[index])]
	}
	return g
}

func (g *passGenerator) dump() string {
	return *(*string)(unsafe.Pointer(&g.password))
}

func Dump() string {
	return initGenerator().sum().dump()
}
