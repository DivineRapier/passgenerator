package storage

import "github.com/DivineRapier/passgenerator/crypto"

type Storage interface {
	Read()
	Write()
	Add(...string)
	Remove(...string)
	Find(string) string
	List() string
	crypto.Crypto
}
