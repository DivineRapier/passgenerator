package passgenerator

import "flag"
import "github.com/DivineRapier/passgenerator/crypto"

var (
	kinds  = flag.String("kind", "1aA", "kinds of letter make up password")
	length = flag.Int("length", 16, "length of password")
	name   = flag.String("name", "", "the username of password")
	get    = flag.String("get", "", "get password")
	add    = flag.String("add", "", "add password with username")
	rm     = flag.String("remove", "", "remove password by username")
	list   = flag.Bool("list", false, "list all username password pairs")
	key    = flag.String("key", crypto.DefaultKey, "key")
)

var (
	digit  = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	lower  = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	upper  = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	symbol = []byte{'~', '!', '@', '#', '$', '%', '^', '*', '{', '}', ',', '?', '(', ')', '_', '+', '-', '='}
)
