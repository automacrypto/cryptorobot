package models

type Config struct {
	Paths struct {
		Macros  string `yaml:"macros"`
		Targets string `yaml:"targets"`
	} `yaml:"paths"`
	Macros []string `yaml:"macros"`
}
