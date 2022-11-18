package apiserver

func (s *APIserver) RegisterHome() {
	s.router.HandleFunc("/", s.HomeHandler()).Methods("GET")
}

func (s *APIserver) RegisterBalance() {
	s.router.HandleFunc("/balance", s.BalanceHandler).Methods("GET", "POST")
}

func (s *APIserver) RegisterAccount() {
	s.router.HandleFunc("/accounts", s.AccountsHandler).Methods("POST")
}

func (s *APIserver) RegisterTransfer() {
	s.router.HandleFunc("/accounts/transfer", s.TransferHandler).Methods("POST")
}

func (s *APIserver) RegisterReport() {
	s.router.HandleFunc("/report", s.ReportHandler).Methods("GET")
}
