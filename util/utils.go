package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func SetData(item DBItem, filePath string) error {
	dbData := getFileData(filePath)
	if keyAlreadyExists(item.Key, dbData) {
    fmt.Println("Key already exists")
    return nil
	}
	dbData = append(dbData, item)
	data, err := json.Marshal(DB{Data: dbData})
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data saved")
	return nil
}

func GetData(path string, key string) (error, DBItem) {

	data := getFileData(path)

	for _, val := range data {
		if val.Key == key {
			return nil, val
		}
	}

	return errors.New("Key not found"), DBItem{}

}

