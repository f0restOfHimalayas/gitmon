package gitmon

type Config struct {
	path []string
}

func LoadConfig() *Config {
	return &Config{}
}
