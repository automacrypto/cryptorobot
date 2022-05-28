package helpers

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func ReadYaml[T any](filename string) T {
	var macro T
	file, read_error := ioutil.ReadFile(filename)
	if read_error != nil {
		panic(read_error)
	}

	unmarshal_error := yaml.Unmarshal(file, &macro)

	if unmarshal_error != nil {
		panic(unmarshal_error)
	}

	return macro
}
