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

// Start initializes and starts the HTTP server
func Start() {
	logger.Info("Starting the HTTP server on localhost:8000")

	// Create a new router using Gorilla Mux
	router := mux.NewRouter()

	// Initialize the PostgreSQL database connection
	dbUser := getPostgresDB()

	// Initialize the UserAuthHandler with its dependencies
	uah := UserAuthHandler{
		services.NewUserAuthService(
			db.NewUserAuthRepositoryDb(dbUser), // User authentication repository
			db.NewUserRepositoryDb(dbUser),     // User repository
		),
	}

	// Initialize the UserHandler with its dependencies
	uh := UserHandler{
		services.NewUserService(db.NewUserRepositoryDb(dbUser)), // User service
	}

	// Define HTTP routes and their corresponding handlers
	router.HandleFunc("/users", uh.IdNo).Methods(http.MethodGet)                              // Get user by ID
	router.HandleFunc("/users/{id_no}", uh.CreateUser).Methods(http.MethodPost)               // Create a new user
	router.HandleFunc("/users/{id_no}", uh.DeleteUser).Methods(http.MethodDelete)             // Delete a user
	router.HandleFunc("/users/{id_no}", uh.UpdateUser).Methods(http.MethodPatch)              // Update user details
	router.HandleFunc("/users/{id_no}/surname", uh.UpdateSurname).Methods(http.MethodPatch)   // Update user surname
	router.HandleFunc("/users/{id_no}/password", uah.CreatePassword).Methods(http.MethodPost) // Create or update user password

	// Configure CORS to allow cross-origin requests
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions}), // Allowed HTTP methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),                                                        // Allowed headers
	)

	logger.Info("HTTP server is ready to accept requests")

	// Start the HTTP server on localhost:8000
	log.Fatal(http.ListenAndServe("localhost:8000", corsHandler(router)))
}

// getPostgresDB establishes a connection to the PostgreSQL database
func getPostgresDB() *sqlx.DB {
	logger.Info("Connecting to PostgreSQL database")

	// Connection string for the PostgreSQL database
	connStr := "user=admin password=Admin123 dbname=email_dir sslmode=disable"

	// Open a new database connection
	userDb, err := sqlx.Open("postgres", connStr)
	if err != nil {
		// Log and terminate the application if the connection fails
		logger.Fatal("Failed to connect to PostgreSQL database", zap.Error(err))
	}

	logger.Info("Successfully connected to PostgreSQL database")
	return userDb
}
