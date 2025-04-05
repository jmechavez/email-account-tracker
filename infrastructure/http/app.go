package http

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmechavez/email-account-tracker/infrastructure/db"
	"github.com/jmechavez/email-account-tracker/infrastructure/logger"
	"github.com/jmechavez/email-account-tracker/internal/ports/services"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func Start() {
	logger.Info("Starting the HTTP server on localhost:8000")

	router := mux.NewRouter()
	dbUser := getPostgresDB()

	uah := UserAuthHandler{
		services.NewUserAuthService(db.NewUserAuthRepositoryDb(dbUser)),
	}

	uh := UserHandler{
		services.NewUserService(db.NewUserRepositoryDb(dbUser)),
	}

	router.HandleFunc("/users", uh.IdNo).Methods(http.MethodGet)
	router.HandleFunc("/users/{id_no}", uh.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id_no}", uh.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id_no}", uh.UpdateUser).Methods(http.MethodPatch)
	router.HandleFunc("/users/{id_no}/surname", uh.UpdateSurname).Methods(http.MethodPatch)
	router.HandleFunc("/users/{id_no}/password", uah.CreatePassword).Methods(http.MethodPost)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	logger.Info("HTTP server is ready to accept requests")
	log.Fatal(http.ListenAndServe("localhost:8000", corsHandler(router)))
}

func getPostgresDB() *sqlx.DB {
	logger.Info("Connecting to PostgreSQL database")
	connStr := "user=admin password=Admin123 dbname=email_dir sslmode=disable"
	userDb, err := sqlx.Open("postgres", connStr)
	if err != nil {
		logger.Fatal("Failed to connect to PostgreSQL database", zap.Error(err))
	}
	logger.Info("Successfully connected to PostgreSQL database")
	return userDb
}
