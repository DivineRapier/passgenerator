package crypto

type Crypto interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
	SetKey(string)
	GetKey() string
}

const DefaultKey = "$}o_L4-J4S(@0)dWag0%#!fW@~31)1F2"
