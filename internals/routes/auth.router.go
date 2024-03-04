package routes

import (
	"bcas/bookstore-go/internals/handlers"
	"bcas/bookstore-go/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitAuthRouter(router *gin.Engine, db *sqlx.DB) {
	// bikin subrouter
	authRouter := router.Group("/auth")
	authRepo := repositories.InitAuthRepo(db)
	authHandler := handlers.InitAuthHandler(authRepo)

	// bikin rute
	//localhost:3000/auth/new
	authRouter.POST("/new", authHandler.Register)
	//localhost:3000/auth
	authRouter.POST("", authHandler.Login)
}
