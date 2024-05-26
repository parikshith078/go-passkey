package util

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
)

func GetDotfilePath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	path := usr.HomeDir + "/go/db/passkey.json"
	return path
}

// BUG: Add create file option if no file exists
func getFileData(path string) []DBItem {
	fileData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var item DB

  if len(fileData) == 0 {
    return item.Data
  }

	if err := json.Unmarshal(fileData, &item); err != nil {
		fmt.Println("Error while unmarshlling")
		panic(err)
	}

	return item.Data

}

func keyAlreadyExists(key string, data []DBItem) bool {
	for _, val := range data {
		if val.Key == key {
			return true
		}
	}
	return false
}
