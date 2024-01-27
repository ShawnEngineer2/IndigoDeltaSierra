package eventemitter

import (
	"encoding/json"
	"os"
)

func EventToFile(vStruct any, outputFilePath string) error {
	//This function writes the passed event struct to the indicated file

	//Create the output file
	f, err := os.Create(outputFilePath)

	if err != nil {
		return err
	}

	defer f.Close()

	//Serialize to JSON, then write to the file
	jsonbytes, err := json.Marshal(vStruct)

	if err != nil {
		return err
	}

	_, err = f.WriteString(string(jsonbytes))

	if err != nil {
		return err
	} else {
		return nil
	}

}
