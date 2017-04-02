package passgenerator

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"flag"
	"fmt"
	"unsafe"
)

const defaultKey = "$}o_L4-J4S(@0)dWag0%#!fW@~31)1F2"

var key = flag.String("key", defaultKey, "key")

func encrypt(plaintext []byte) []byte {
	block, err := aes.NewCipher(*(*[]byte)(unsafe.Pointer(key)))
	if err != nil {
		panic(err)
	}

	nonce, _ := hex.DecodeString("37b8e8a308c354048d245f6d")
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return []byte(fmt.Sprintf("%x", ciphertext))
}

func decrypt(ciphertext []byte) []byte {
	nonce, _ := hex.DecodeString("37b8e8a308c354048d245f6d")
	ciphertext, _ = hex.DecodeString(*(*string)(unsafe.Pointer(&ciphertext)))

	block, err := aes.NewCipher(*(*[]byte)(unsafe.Pointer(key)))
	if err != nil {
		panic(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	return plaintext
}
