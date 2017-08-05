package util

import "io/ioutil"

// StringToFile stores a string in a file on the disk
func StringToFile(file, data string) error {
	// TODO: add testcase for failing WriteFile
	err := ioutil.WriteFile(file, []byte(data), 0666)
	if err != nil {
		return err
	}
	return nil
}

// FileToString reads the content of a file
func FileToString(file string) (string, error) {
	src, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(src), nil
}
