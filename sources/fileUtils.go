package sources

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func RemoveContents(dirPath string) error {
	directory, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer directory.Close()
	names, err := directory.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dirPath, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteFileOnServer(pathToFolder string, prefix string, extension string, data []byte) string {
	outputFile, err := ioutil.TempFile(pathToFolder, prefix+"*"+extension)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		return ""
	}
	defer outputFile.Close()

	outputFile.Write(data)
	return outputFile.Name()
}

func CreateUuidKey() []byte {
	uuidWithHyphens := uuid.New()
	uuidWithoutHyphens := strings.Replace(uuidWithHyphens.String(), "-", "", -1)
	return []byte(uuidWithoutHyphens)
}
