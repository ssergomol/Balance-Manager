package main

import (
	"github.com/sirupsen/logrus"
	"github.com/ssergomol/Balance-Manager/pkg/apiserver"
)

func main() {
	config := apiserver.NewConfig()
	server := apiserver.CreateServer(config)

	if err := server.Start(); err != nil {
		logrus.Fatal(err)
	}
}
