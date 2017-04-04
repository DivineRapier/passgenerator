package passgenerator

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unsafe"
)

type passGenerator struct {
	container [][]byte
	length    int
	password  []byte
}

var (
	digit  = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	lower  = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	upper  = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	symbol = []byte{'~', '!', '@', '#', '$', '%', '^', '*', '{', '}', ',', '?', '(', ')', '_', '+', '-', '='}
)

var (
	kinds  = flag.String("kind", "1aA", "kinds of letter make up password")
	length = flag.Int("length", 16, "length of password")
	name   = flag.String("name", "", "the username of password")
	get    = flag.String("get", "", "get password")
	add    = flag.String("add", "", "add password with username")
	rm     = flag.String("remove", "", "remove password by username")
)

var generator passGenerator

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.Usage = help
	flag.Parse()
	initGenerator()
}

func initGenerator() {
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
	return
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

// Dump dump the password
func dump() string {
	pwd := generator.sum().dump()
	user[*name] = pwd
	return pwd
}

// Run run app
func Run() {
	if len(*name) > 0 && len(*get) > 0 {
		fmt.Println("You may specify one and only on of '-name', or '-get' option")
		return
	}

	switch {
	case *list:
		show()
	case len(*rm) > 0:
		remove(*rm)
		write()
	case len(*add) > 0:
		addPassword()
		write()
	case len(*name) > 0:
		dump()
		write()
	case len(*get) > 0:
		fmt.Println(find(*get))
	default:
		dump()
	}
}

func help() {
	fmt.Println(
		`
Usage: 
	
	./generator -option  arguments

The options are:

	kind            The password contains an instance of some kind of character
	length          The length of the password
	name            Set the password for the specified account
	get             Get the password for the specified account
	list            List all passwords
	remove          Remove password by username
	add             Add pair of username and password (user:password)
`,
	)
}

func addPassword() {
	index := strings.IndexByte(*add, ':')
	if index <= 0 || index == len(*add)-1 {
		fmt.Println("Error: invalid format. eg user:password")
		os.Exit(-1)
	}

	username := (*add)[:index]
	password := (*add)[index+1:]

	user[username] = password
}
