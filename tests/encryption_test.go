package tests

import (
	"bytes"
	"file-encryptor/sources"
	"io/ioutil"
	"os"
	"testing"
)

func TestEncryption(t *testing.T) {
	fileUrl := "../test-files/download.jpg"
	file, err := os.Open(string(fileUrl))
	if err != nil {
		t.Fatalf("Could not open file: %v\n", err)
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Could not slice file to bytes: %v\n", err)
	}

	key := sources.CreateUuidKey()

	encryptedBytes := sources.EncryptBytes(fileBytes, key)
	decryptedBytes := sources.DecryptBytes(encryptedBytes, key)

	res := bytes.Compare(fileBytes, decryptedBytes)

	if res != 0 {
		t.Fatalf("Encryption algorithm not working")
	}
}
