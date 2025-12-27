package config

type WebServerConfig struct {
	Port uint16 `yaml:"port"`

	ServerName string `yaml:"server_name"`
	Locale     string `yaml:"locale"` // e.g., "ko-KR", "en-US"

	MongoUri       string `yaml:"mongo_uri"`
	MongoEnvDBName string `yaml:"mongo_env_db_name"`
}
