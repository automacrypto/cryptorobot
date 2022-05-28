package models

type Macro struct {
	Version     string `yaml:"version"`
	Application string `yaml:"application"`
	Name        string `yaml:"name"`
	Enabled     bool   `yaml:"enabled"`
	Steps       []struct {
		Name    string `yaml:"name"`
		Enabled bool   `yaml:"enabled"`
		Event   struct {
			Type  string `yaml:"type"`
			Value string `yaml:"value"`
		} `yaml:"event"`
	} `yaml:"steps"`
}
