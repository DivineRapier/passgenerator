package passgenerator

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.Usage = help
	flag.Parse()
}

// Run run app
func Run() {
	generator := NewPassGenerator()

	if len(*name) > 0 && len(*get) > 0 {
		fmt.Println("You may specify one and only on of '-name', or '-get' option")
		return
	}

	switch {
	case *list:
		fmt.Println(generator.List())
	case len(*rm) > 0:
		rmList := strings.Split(*rm, " ")
		generator.Remove(rmList...)
		generator.Write()
	case len(*add) > 0:
		addList := strings.Split(*add, " ")
		generator.Add(addList...)
		generator.Write()
	case len(*name) > 0:
		tmp := generator.Sum().Dump()
		fmt.Println(tmp)
		generator.Add(*name, tmp)
		generator.Write()
	case len(*get) > 0:
		fmt.Println(generator.Find(*get))
	default:
		fmt.Println(generator.Dump())
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
