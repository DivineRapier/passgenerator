package passgenerator

import (
	"flag"
	"fmt"

	"github.com/apcera/termtables"
)

var list = flag.Bool("list", false, "list all username password pairs")

func show() {
	table := termtables.CreateTable()

	table.AddHeaders("username", "password")
	for username, password := range user {
		table.AddRow(username, password)
	}
	fmt.Println(table.Render())
}
