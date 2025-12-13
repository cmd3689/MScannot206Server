package config

type WebClientConfig struct {
	Url  string `yaml:"url"`
	Port uint16 `yaml:"port"`
}
