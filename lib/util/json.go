package util

import (
	"bytes"
	"encoding/json"
)

// InterfaceToJSON returns a byte slice with the json data
func InterfaceToJSON(data interface{}) ([]byte, error) {
	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return j, nil
}

// InterfaceToJSONString returns a string with the json data
func InterfaceToJSONString(data interface{}) (string, error) {
	j, err := InterfaceToJSON(data)
	return string(j), err
}

// InterfaceToJSONStringPretty returns a pretty version of the json string
func InterfaceToJSONStringPretty(data interface{}) (string, error) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "    ")
	err := enc.Encode(data)

	return buf.String(), err
}
