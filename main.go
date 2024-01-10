package main

import (
	"theater/config"
	"theater/handlers"
	"theater/movie"
	"theater/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadAppConfig()
	db, err := config.Connect()

	if err != nil {
		panic(err.Error())
	}

	movieRepository := movie.NewRepository(db)

	movieService := movie.NewService(movieRepository)

	movieHandler := handlers.NewMovieHandler(movieService)

	router := gin.Default()

	movieAPI := router.Group("/api/v1/movie")

	routes.MovieRoutes(movieAPI, movieHandler)

	router.Run()
}
