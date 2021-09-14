package tests

import (
	"file-encryptor/sources"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteFileOnServerValidFile(t *testing.T) {
	fileUrl := "../test-files/download.jpg"
	file, err := os.Open(string(fileUrl))
	if err != nil {
		t.Fatalf("Could not open file: %v\n", err)
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Could not slice file to bytes: %v\n", err)
	}

	outputFilePath := sources.WriteFileOnServer("../test-files/", "test-", ".txt", fileBytes)

	//Check if file does not exist
	if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
		t.Fatalf("File not created successfully")
	}

	//Remove the file
	err = os.Remove(outputFilePath)
	if err != nil {
		t.Errorf("Cannot delete existing file")
	}
}

func TestWriteFileOnServerInvalidFile(t *testing.T) {
	//Empty byte array
	var fileBytes []byte

	outputFilePath := sources.WriteFileOnServer("../test-files/", "test-", ".txt", fileBytes)

	//Check if file does not exist
	if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
		t.Fatalf("File not created successfully")
	}

	//Remove the file
	err := os.Remove(outputFilePath)
	if err != nil {
		t.Errorf("Cannot delete existing file")
	}
}

func TestRemoveContentsWithNonEmptyFolder(t *testing.T) {
	fileUrl := "../test-files/download.jpg"
	file, err := os.Open(string(fileUrl))
	if err != nil {
		t.Fatalf("Could not open file: %v\n", err)
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Could not slice file to bytes: %v\n", err)
	}

	//Create 3 files with different names
	//1
	outputFile1 := sources.WriteFileOnServer("../test-files/remove-all/", "test-", ".txt", fileBytes)

	//In case file does not exist
	if _, err := os.Stat(outputFile1); os.IsNotExist(err) {
		t.Fatalf("File1 is not created successfully")
	}
	//2
	outputFile2 := sources.WriteFileOnServer("../test-files/remove-all/", "test-", ".txt", fileBytes)

	//In case file does not exist
	if _, err := os.Stat(outputFile2); os.IsNotExist(err) {
		t.Fatalf("File2 is not created successfully")
	}
	//3
	outputFile3 := sources.WriteFileOnServer("../test-files/remove-all/", "test-", ".txt", fileBytes)

	//In case file does not exist
	if _, err := os.Stat(outputFile3); os.IsNotExist(err) {
		t.Fatalf("File3 is not created successfully")
	}

	//Here we know that we have some files in folder test-files
	sources.RemoveContents("../test-files/remove-all/")

	//Check that files are removed
	//1
	if _, err := os.Stat("../test-files/remove-all/" + outputFile1); err == nil {
		t.Fatalf("File is not removed successfully")
	}
	//2
	if _, err := os.Stat("../test-files/remove-all/" + outputFile2); err == nil {
		t.Fatalf("File is not removed successfully")
	}
	//3
	if _, err := os.Stat("../test-files/remove-all/" + outputFile3); err == nil {
		t.Fatalf("File is not removed successfully")
	}
}

func TestRemoveContentsWithEmptyFolder(t *testing.T) {
	//Check that folder is empty
	res, err := IsEmptyFolder("../test-files/remove-all/")
	if err != nil && !res {
		t.Fatalf("Folder is not empty")
	}

	//Here we know that we do not have some files in folder test-files
	sources.RemoveContents("../test-files/remove-all/")

	//Check that folder is empty
	res, err = IsEmptyFolder("../test-files/remove-all/")
	if err != nil && !res {
		t.Fatalf("Folder is not empty")
	}
}

func TestRemoveContentsWithNonExistingFolder(t *testing.T) {
	//Check that folder does not exist
	if _, err := os.Stat("../test-files/remove-all-1/"); err == nil {
		t.Fatalf("Folder exists")
	}

	//Here we know that we do not have some files in folder test-files
	sources.RemoveContents("../test-files/remove-all-1/")
}

func IsEmptyFolder(dirPath string) (bool, error) {
	f, err := os.Open(dirPath)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
