package apiserver

type ConfigServer struct {
	BindAddress string
	DatabaseURL string
	LogLevel    string
}

func NewConfig() *ConfigServer {
	return &ConfigServer{
		BindAddress: ":8080",
		LogLevel:    "debug",
	}
}
