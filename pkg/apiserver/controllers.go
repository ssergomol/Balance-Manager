package apiserver

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ssergomol/Balance-Manager/pkg/models"
)

func (s *APIserver) HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello! This is balance manager</h1>"))
	}
}

func (s *APIserver) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.logger.Info("got balance POST request")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			s.logger.Fatal(err)
		}

		balance := models.Balance{}

		err = json.Unmarshal([]byte(body), &balance)
		if err != nil {
			s.logger.Fatal(err)
		}

		err = s.db.Balance().ReplenishBalance(balance)
		if err != nil {
			s.logger.Fatal(err)
		}

		bytes, err := json.Marshal(balance)
		if err != nil {
			s.logger.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)

	case http.MethodGet:
		s.logger.Info("got balance get request")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			s.logger.Fatal(err)
		}

		balance := models.Balance{}

		err = json.Unmarshal([]byte(body), &balance)
		if err != nil {
			s.logger.Fatal(err)
		}

		getBalance, err := s.db.Balance().GetBalance(balance.ID)
		if err != nil {
			s.logger.Fatal(err)
		}

		bytes, err := json.Marshal(getBalance)
		if err != nil {
			s.logger.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

func (s *APIserver) AccountsHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("got accounts POST request")
	account := models.Account{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		s.logger.Fatal(err)
	}

	err = json.Unmarshal([]byte(body), &account)
	if err != nil {
		s.logger.Fatal(err)
	}

	switch account.ServiceID {
	// Debit funds from the balance
	case 1:
		balance, err := s.db.Balance().GetBalance(account.UserID)
		if err != nil {
			s.logger.Fatal(err)
		}

		if balance.Sum < account.Sum {
			w.WriteHeader(http.StatusBadRequest)
			message, err := json.Marshal("Not enough funds on the balance")
			if err != nil {
				s.logger.Fatal(err)
			}

			w.Write(message)
			return
		}

	// Debit funds from the account
	case 2:
		acc, err := s.db.Account().GetAccountSum(account.ID, account.UserID)
		if err != nil {
			s.logger.Fatal(err)
		}

		if acc.Sum < account.Sum {
			w.WriteHeader(http.StatusBadRequest)
			message, err := json.Marshal("Not enough funds on the account")
			if err != nil {
				s.logger.Fatal(err)
			}

			w.Write(message)
			return
		}

	}

	s.db.Account().ReserveFunds(account)

	if err != nil {
		s.logger.Fatal(err)
	}

	bytes, err := json.Marshal(account)
	if err != nil {
		s.logger.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
