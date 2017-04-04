package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"unsafe"
)

const nonceString = "37b8e8a308c354048d245f6d"

type AES struct {
	key string
}

func NewAES(key string) Crypto {
	ret := new(AES)
	ret.key = key
	return ret
}

func (a *AES) Encrypt(plaintext []byte) []byte {
	block, err := aes.NewCipher(*(*[]byte)(unsafe.Pointer(&a.key)))
	if err != nil {
		panic(err)
	}

	nonce, _ := hex.DecodeString(nonceString)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return []byte(fmt.Sprintf("%x", ciphertext))
}

func (a *AES) Decrypt(ciphertext []byte) []byte {
	nonce, _ := hex.DecodeString(nonceString)
	ciphertext, _ = hex.DecodeString(*(*string)(unsafe.Pointer(&ciphertext)))

	block, err := aes.NewCipher(*(*[]byte)(unsafe.Pointer(&a.key)))
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

func (a *AES) SetKey(key string) {
	a.key = key
}

func (a *AES) GetKey() string {
	return a.key
}
