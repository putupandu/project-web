// File: backend/cmd/api/main.go
package main

import (
	"e-library/backend/internal/config"
	"e-library/backend/internal/database"
	"e-library/backend/internal/handlers"
	"e-library/backend/internal/middleware"
	"e-library/backend/internal/repositories"
	"e-library/backend/internal/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	log.Println("Configuration loaded successfully")

	// Initialize database
	db := database.InitDB(cfg)
	defer db.Close()
	log.Println("Database connected successfully")

	// Initialize repositories
	bookRepo := repositories.NewBookRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)

	// Initialize services
	bookService := services.NewBookService(bookRepo, categoryRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	searchService := services.NewSearchService(bookRepo)

	// Initialize handlers
	bookHandler := handlers.NewBookHandler(bookService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	searchHandler := handlers.NewSearchHandler(searchService)
	healthHandler := handlers.NewHealthHandler()

	// Setup router
	router := mux.NewRouter()

	// Apply middleware
	router.Use(middleware.Logger)
	router.Use(middleware.CORS)

	// API prefix
	api := router.PathPrefix("/api").Subrouter()

	// Root test endpoint
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API is running"))
	}).Methods("GET")

	// Health check
	api.HandleFunc("/health", healthHandler.HealthCheck).Methods("GET")

	// Book routes
	api.HandleFunc("/books", bookHandler.GetAllBooks).Methods("GET")
	api.HandleFunc("/books/{id}", bookHandler.GetBookByID).Methods("GET")
	api.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	api.HandleFunc("/books/{id}", bookHandler.UpdateBook).Methods("PUT")
	api.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods("DELETE")
	api.HandleFunc("/books/{id}/download", bookHandler.IncrementDownload).Methods("POST")
	api.HandleFunc("/books/{id}/view", bookHandler.IncrementView).Methods("POST")

	// Category routes
	api.HandleFunc("/categories", categoryHandler.GetAllCategories).Methods("GET")
	api.HandleFunc("/categories/{id}", categoryHandler.GetCategoryByID).Methods("GET")

	// Search route
	api.HandleFunc("/search", searchHandler.Search).Methods("GET")

	// Serve static files (uploads)
	router.PathPrefix("/uploads/").Handler(
		http.StripPrefix("/uploads/",
			http.FileServer(http.Dir("./uploads"))))

	// Start server
	addr := ":" + cfg.Port
	log.Printf("Server starting on port %s...", cfg.Port)
	log.Fatal(http.ListenAndServe(addr, router))
}
