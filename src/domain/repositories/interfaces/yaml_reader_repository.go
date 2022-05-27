package domain

type IYamlReaderRepository interface {
	Read() error
}
