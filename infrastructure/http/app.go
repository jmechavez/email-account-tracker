package http

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmechavez/email-account-tracker/infrastructure/db"
	"github.com/jmechavez/email-account-tracker/internal/ports/services"
	"github.com/jmoiron/sqlx"
)

func Start() {
	router := mux.NewRouter()

	dbUser := getPostgresDB()

	//UserRepositorydb := db.NewUserRepositoryDb(dbUser)

	uh := UserHandler{
		services.NewUserService(db.NewUserRepositoryDb(dbUser)),
	}

	router.HandleFunc("/users", uh.IdNo).Methods(http.MethodGet)
	router.HandleFunc("/users/{id_no}", uh.CreateUser).Methods(http.MethodPost)
	// Set CORS options
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Change this to specific origins if needed
		handlers.AllowedMethods(
			[]string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
				http.MethodOptions,
			},
		),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Fatal(http.ListenAndServe("localhost:8000", corsHandler(router)))
}

func getPostgresDB() *sqlx.DB {
	connStr := "user=admin password=Admin123 dbname=email_dir sslmode=disable"
	userDb, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return userDb
}
