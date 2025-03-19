package main

import (
	"github.com/jmechavez/email-account-tracker/infrastructure/http"
	"github.com/jmechavez/email-account-tracker/infrastructure/logger"
)

func main() {
	logger.Info("Starting the application")
	http.Start()
}
