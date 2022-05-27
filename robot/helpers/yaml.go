package helpers

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func ReadYaml(filename string) []byte {
	file, read_error := ioutil.ReadFile(filename)

	if read_error != nil {
		panic(read_error)
	}

	data := make(map[interface{}]interface{})

	unmarshal_error := yaml.Unmarshal(file, &data)

	if unmarshal_error != nil {
		panic(unmarshal_error)
	}

	return file
}
