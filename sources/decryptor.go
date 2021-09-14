package sources

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func DecryptBytes(cipherBytes []byte, key []byte) []byte {
	//key := sha256.Sum256([]byte("kitty"))

	//generate a new aes cipher using 32 byte long key
	block, err := aes.NewCipher(key)
	//handle errors
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(block)
	//handle errors
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(cipherBytes) < nonceSize {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}

	nonce, ciphertext := cipherBytes[:nonceSize], cipherBytes[nonceSize:]
	decryptedBytes, err := gcm.Open(nil, nonce, ciphertext, nil)
	//Handle error
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}

	return decryptedBytes
	//---------------------------
}
