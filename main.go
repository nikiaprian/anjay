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

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://fe.codein.studio"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/auth/check-token", handler.CheckToken)
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
	router.GET("/blogs/:id", handler.GetBlogByID)

	router.POST("/like/forum/:id", handler.CheckUserRole, handler.CreateLikeByForumId)
	router.DELETE("/like/forum/:id", handler.CheckUserRole, handler.DeleteLikeByForumId)

	router.GET("/forums", handler.CheckUserLoginOptional, handler.GetAllForum)
	router.GET("/forums/:id", handler.CheckUserLoginOptional, handler.GetForumById)
	router.POST("/forums/new", handler.CheckUserRole, handler.CreateForum)
	// router.PUT("/forums/:id", handler.CheckUserRole, handler.UpdateForum)
	router.DELETE("/forums/:id", handler.CheckUserRole, handler.DeleteForum)

	router.PATCH("/forum/comment/:id/selected-answer", handler.CheckUserRole, handler.SelectedCommentAnswer)

	router.POST("/commentsforum/:id", handler.CheckUserRole, handler.CreateCommentForum)
	router.GET("/commentsforum/:id", handler.GetAllCommentByForumID)
	router.DELETE("/commentsforum/:id", handler.CheckUserRole, handler.DeleteCommentForum)

	router.POST("/comments/:id", handler.CheckUserRole, handler.CreateCommentBlog)
	router.GET("/comments/:id", handler.GetAllCommentByBlogID)
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
