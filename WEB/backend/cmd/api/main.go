//main
package main

import (
	"log"
	"net/http"

	"e-library/backend/internal/config"
	"e-library/backend/internal/database"
	"e-library/backend/internal/handlers"
	"e-library/backend/internal/middleware"
	"e-library/backend/internal/repositories"
	"e-library/backend/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()
	db := database.InitDB(cfg)
	defer db.Close()

	bookRepo := repositories.NewBookRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)
	userRepo := repositories.NewUserRepository(db)
	savedBookRepo := repositories.NewSavedBookRepository(db)

	bookService := services.NewBookService(bookRepo, categoryRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	searchService := services.NewSearchService(bookRepo)
	userService := services.NewUserService(userRepo)
	savedBookService := services.NewSavedBookService(savedBookRepo)
	authService := services.NewAuthService(userRepo)

	bookHandler := handlers.NewBookHandler(bookService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	searchHandler := handlers.NewSearchHandler(searchService)
	userHandler := handlers.NewUserHandler(userService)
	savedBookHandler := handlers.NewSavedBookHandler(savedBookService)
	authHandler := handlers.NewAuthHandler(authService)
	healthHandler := handlers.NewHealthHandler()

	r := mux.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CORS)

	api := r.PathPrefix("/api").Subrouter()

	// Auth
	api.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	api.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	// Books
	api.HandleFunc("/books", bookHandler.GetAllBooks).Methods("GET")
	api.HandleFunc("/books/{id}", bookHandler.GetBookByID).Methods("GET")
	api.HandleFunc("/books/{id}/view", bookHandler.GetBookByID).Methods("GET", "POST") // ← BARU: tambahan untuk /view

	// Categories
	api.HandleFunc("/categories", categoryHandler.GetAllCategories).Methods("GET")

	// Search
	api.HandleFunc("/search", searchHandler.Search).Methods("GET")

	// Users (CRUD)
	api.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Health check
	api.HandleFunc("/health", healthHandler.HealthCheck).Methods("GET")

	// Secured routes (require auth)
	secured := api.NewRoute().Subrouter()
	secured.Use(middleware.Auth)

	secured.HandleFunc("/saved-books", savedBookHandler.SaveBook).Methods("POST")
	secured.HandleFunc("/saved-books", savedBookHandler.GetSavedBooks).Methods("GET")
	secured.HandleFunc("/saved-books/{id}", savedBookHandler.RemoveSavedBook).Methods("DELETE")

	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))

	log.Println("Server running on :" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}// register http routes
