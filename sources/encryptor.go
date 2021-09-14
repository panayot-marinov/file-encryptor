package sources

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func EncryptBytes(fileBytes []byte, key []byte) []byte {
	//generate a new aes cipher using our 32 byte long key
	block, err := aes.NewCipher(key)
	//handle errors
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(block)
	//handle errors
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	//---------------------------

	return gcm.Seal(nonce, nonce, fileBytes, nil)
}
