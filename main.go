package main

import (
	"kel15/config"
	"kel15/handler"
	"kel15/project"
	"kel15/repository"
	"kel15/storage"
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
	storage := storage.NewS3()

	newProject := project.NewProject(usecase, *storage)

	handler := handler.NewHandler(newProject)

	router := gin.Default()

	router.POST("/auth/login", handler.UserLogin)
	router.POST("/auth/register", handler.UserRegister)
	router.GET("/auth/login/:provider", handler.UserLoginByProvider)
	router.GET("/testing-middleware-user", handler.CheckUserRole, handler.TestingMiddlewareUser)
	router.GET("/testing-middleware-admin", handler.CheckAdminRole, handler.TestingMiddlewareAdmin)
	router.GET("/auth/callback/:provider", handler.UserLoginByProviderCallback)

	router.GET("/users", handler.UserList)
	router.POST("/users/:id", handler.UserProfileUpdate)

	router.GET("/blogs", handler.GetAllBlog)
	router.POST("/blogs/new", handler.CheckUserRole, handler.CreateBlog)
	// router.PUT("/blogs/:id", handler.CheckUserRole, handler.UpdateBlog)
	router.DELETE("/blogs/:id", handler.CheckUserRole, handler.DeleteBlog)
	router.GET("/blogs/:id", handler.GetBlog)

	router.GET("/forums", handler.GetAllForum)
	router.POST("/forums/new", handler.CreateForum)
	router.PUT("/forums/:id", handler.UpdateForum)
	router.DELETE("/forums/:id", handler.DeleteForum)
	router.GET("/forums/:id", handler.GetForum)

	router.POST("/commentsforum/:id", handler.CheckUserRole, handler.CreateCommentForum)
	router.GET("/commentsforum/:id", handler.CheckUserRole, handler.GetAllCommentByForumID)
	router.DELETE("/commentsforum/:id", handler.CheckUserRole, handler.DeleteCommentForum)

	router.POST("/comments/:id", handler.CheckUserRole, handler.CreateCommentBlog)
	router.GET("/comments/:id", handler.CheckUserRole, handler.GetAllCommentByBlogID)
	router.DELETE("/comments/:id", handler.CheckUserRole, handler.DeleteCommentByID)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:9090",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
