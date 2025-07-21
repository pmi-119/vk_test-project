package main

import (
	"log"
	"net/http"

	add_product_handler "VK_test_proect/internal/api/add_product"
	authorize_handler "VK_test_proect/internal/api/authorize"
	get_feed_handler "VK_test_proect/internal/api/get_feed"
	register_handler "VK_test_proect/internal/api/register"

	"VK_test_proect/internal/postgres"
	product_repo "VK_test_proect/internal/repository/product_info"
	user_repo "VK_test_proect/internal/repository/user_info"
	"VK_test_proect/internal/service/add_product"
	"VK_test_proect/internal/service/authorize"
	"VK_test_proect/internal/service/get_feed"
	"VK_test_proect/internal/service/register"
	tokenservice "VK_test_proect/internal/service/tokenService"

	"github.com/go-chi/chi"
)

func main() {
	// DB

	db, err := postgres.Connect()
	if err != nil {
		panic(err)
	}

	// Migrations

	migrator := postgres.NewMigrator(db)
	if err := migrator.Init(); err != nil {
		panic(err)
	}

	// Repository

	userRepo := user_repo.New(db)
	productRepo := product_repo.New(db)

	// Services

	tokenService := tokenservice.New("secret")
	addProductService := add_product.New(productRepo, userRepo)
	authorizeService := authorize.New(tokenService, userRepo)
	getFeedService := get_feed.New(productRepo)
	registerService := register.New(userRepo)

	// Handlers

	addProductHandler := add_product_handler.New(addProductService, tokenService)
	authorizeHandler := authorize_handler.New(authorizeService)
	getfeedHandler := get_feed_handler.New(getFeedService, tokenService)
	registerHandler := register_handler.New(registerService)

	router := chi.NewRouter()

	// Routes
	router.Post("/register", registerHandler.PostRegister)
	router.Post("/authorize", authorizeHandler.PostLogin)
	router.Post("/add", addProductHandler.PostFeed)
	router.Get("/feed", getfeedHandler.PostFeed)

	log.Println("Starting server on :8080...")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
