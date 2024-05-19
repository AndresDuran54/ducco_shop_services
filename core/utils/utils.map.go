package utils

import (
	"encoding/json"
)

type UtilsMap struct {
}

func (o UtilsMap) InterfaceToStruct(in interface{}, out interface{}) error {
	var err error
	var data []byte

	if data, err = json.Marshal(in); err != nil {
		return err
	}
	err = json.Unmarshal(data, &out)
	return err
}

func (o UtilsMap) InterfaceToMap(in interface{}) (string, error) {
	data, err := json.Marshal(in)

	if err != nil {
		return "", err
	}

	dataStr := string(data)

	return dataStr, nil
}
