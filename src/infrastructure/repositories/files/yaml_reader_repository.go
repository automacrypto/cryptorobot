package infrastructure

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type YamlReaderRepository struct {
}

func (r YamlReaderRepository) Read(filename string) {
	file, read_error := ioutil.ReadFile(filename)

	if read_error != nil {
		panic(read_error)
	}

	data := make(map[interface{}]interface{})

	unmarshal_error := yaml.Unmarshal(file, &data)

	if unmarshal_error != nil {
		panic(unmarshal_error)
	}

}
