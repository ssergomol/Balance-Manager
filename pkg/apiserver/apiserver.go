package apiserver

type APIserver struct {
	config *ConfigServer
}

func CreateServer(config *ConfigServer) *APIserver {
	return &APIserver{
		config: config,
	}
}

func (s *APIserver) Start() error {

}
