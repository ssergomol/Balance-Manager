package apiserver

func (s *APIserver) RegisterHome() {
	s.router.HandleFunc("/", s.HomeHandler()).Methods("GET")
}
