package apiserver

type Config struct {
	BindAdress string `toml:"bind_adress"`
	LogLevel   string `toml: "log_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAdress: ":8080",
		LogLevel:   "debug",
	}
}
