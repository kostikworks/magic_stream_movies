package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kostikworks/magic_stream_movies/server/magic_stream_movies_server/controllers"
	"github.com/kostikworks/magic_stream_movies/server/magic_stream_movies_server/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupProtectedRoutes(router *gin.Engine, client *mongo.Client) {
	router.Use(middleware.AuthMiddleware())

	router.GET("/movie/:imdb_id", controller.GetMovie(client))
	router.POST("/addmovie", controller.AddMovie(client))
	router.GET("/recommendedmovies", controller.GetRecommendedMovies(client))
	router.PATCH("/updatereview/:imdb_id", controller.AdminReviewUpdate(client))
}
