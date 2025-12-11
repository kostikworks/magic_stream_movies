package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kostikworks/magic_stream_movies/server/magic_stream_movies_server/controllers"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupUnprotectedRoutes(router *gin.Engine, client *mongo.Client) {

	router.GET("/movies", controller.GetMovies(client))
	router.POST("/register", controller.RegisterUser(client))
	router.POST("login", controller.LoginUser(client))
	router.GET("/genres", controller.GetGenres(client))
	router.GET("/refresh", controller.RefreshTokenHandler(client))
}
