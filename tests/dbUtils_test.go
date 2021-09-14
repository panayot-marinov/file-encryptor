package tests

import (
	"bytes"
	"file-encryptor/sources"
	"testing"

	_ "github.com/lib/pq" //postgres drivers for initialization
)

//TODO:: mock table connection
func TestGetKeyValidId(t *testing.T) {
	db := sources.ConnectToDb()
	defer db.Close()

	fileId := "267027311"
	value, err := sources.GetKey(db, fileId)
	if err != nil {
		t.Errorf("Cannot get key")
	}

	valueExpected := []byte("35cc4e0b1f5942b8be5ec3413ff80fe1")

	res := bytes.Compare(value, valueExpected)

	if res != 0 {
		t.Fatalf("Keys are not equal")
	}
}

func TestGetKeyInvalidId(t *testing.T) {
	db := sources.ConnectToDb()
	defer db.Close()

	fileId := "267027312"
	_, err := sources.GetKey(db, fileId)
	if err == nil {
		t.Errorf("Got value with invalid id")
	}
}

func TestGetFilePath(t *testing.T) {
	db := sources.ConnectToDb()
	defer db.Close()

	fileId := "267027311"
	value, err := sources.GetFilePath(db, fileId)
	if err != nil {
		t.Errorf("Cannot get key")
	}

	valueExpected := []byte("uploaded-images\\upload-267027311.jpg")

	res := bytes.Compare(value, valueExpected)

	if res != 0 {
		t.Fatalf("File paths are not equal")
	}
}

func TestGetFilePathInvalidId(t *testing.T) {
	db := sources.ConnectToDb()
	defer db.Close()

	fileId := "267027312"
	_, err := sources.GetKey(db, fileId)
	if err == nil {
		t.Errorf("Got value with invalid id")
	}
}

func TestGetOrigFileName(t *testing.T) {
	db := sources.ConnectToDb()
	defer db.Close()

	fileId := "267027311"
	value, err := sources.GetOrigFileName(db, fileId)
	if err != nil {
		t.Errorf("Cannot get key")
	}

	valueExpected := []byte("Sample_abc.jpg")

	res := bytes.Compare(value, valueExpected)

	if res != 0 {
		t.Fatalf("Orig file names are not equal")
	}
}

func TestGetOrigFileNameInvalidId(t *testing.T) {
	db := sources.ConnectToDb()
	defer db.Close()

	fileId := "267027312"
	_, err := sources.GetKey(db, fileId)
	if err == nil {
		t.Errorf("Got value with invalid id")
	}
}
