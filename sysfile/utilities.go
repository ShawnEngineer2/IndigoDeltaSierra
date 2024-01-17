package sysfile

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadFileToStruct(filepath string, targetStruct any) bool {

	sourceFile := OpenFile(filepath)

	if sourceFile == nil {
		fmt.Println("Cannot open config file ... startup terminated")
		return false
	}

	defer CloseFile(sourceFile)

	fileBytes, _ := io.ReadAll(sourceFile)

	err := json.Unmarshal(fileBytes, targetStruct)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true

}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DirExists(dirpath string) bool {
	info, err := os.Stat(dirpath)

	if err != nil {
		fmt.Println(err.Error())
	}

	return info.IsDir()
}

func OpenFile(filename string) *os.File {

	//Attempt to open passed file, check for error
	targetFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return targetFile

}

func CloseFile(filepointer *os.File) bool {
	//Return True if the file was closed - False otherwise
	err := filepointer.Close()

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
