package apiserver

import "net/http"

func (s *APIserver) HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello! This is balance manager</h1>"))
	}
}
