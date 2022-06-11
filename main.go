package main

import (
	"kel15/config"
	"kel15/handler"
	"kel15/project"
	"kel15/repository"
	"kel15/usecase"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	db, err := config.NewSQLLiteDB()
	if err != nil {
		panic(err)
	}

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(&repository)

	newProject := project.Project{
		Usecase: usecase,
	}
	handler := handler.NewHandler(&newProject)

	router := gin.Default()

	router.POST("/auth/login", handler.UserLogin)
	router.POST("/auth/register", handler.UserRegister)
	router.GET("/auth/login/:provider", handler.UserLoginByProvider)
	router.GET("/auth/callback/:provider", handler.UserLoginByProviderCallback)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:9090",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
