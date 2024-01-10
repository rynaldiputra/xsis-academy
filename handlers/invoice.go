package handlers

import (
	"net/http"
	"strconv"
	"theater/helper"
	"theater/movie"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieService movie.Service
}

func NewMovieHandler(movieService movie.Service) *MovieHandler {
	return &MovieHandler{movieService}
}

func (mb *MovieHandler) Get(c *gin.Context) {
	movie, err := mb.movieService.GetMovies()

	if err != nil {
		response := helper.JsonResponse("Gagal mendapatkan movie", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if len(movie) == 0 {
		response := helper.JsonResponse("Data movie kosong", http.StatusOK, "success", movie)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.JsonResponse("movie berhasil diambil", http.StatusOK, "success", movie)

	c.JSON(http.StatusOK, response)
}

func (mb *MovieHandler) Find(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	movie, err := mb.movieService.GetMovie(id)

	if err != nil {
		response := helper.JsonResponse("Gagal mendapatkan movie", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("movie berhasil diambil", http.StatusOK, "success", movie)

	c.JSON(http.StatusOK, response)
}

func (mh *MovieHandler) Create(c *gin.Context) {
	var input movie.MovieInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal menyimpan movie", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newMovie, err := mh.movieService.CreateMovie(input)

	if err != nil {
		response := helper.JsonResponse("Gagal menyimpan movie", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := movie.FormatMovie(newMovie)
	response := helper.JsonResponse("Penyimpanan movie berhasil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}

func (mh *MovieHandler) Edit(c *gin.Context) {
	var input movie.MovieInput
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal menyimpan movie", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newMovie, err := mh.movieService.EditMovie(input, id)

	if err != nil {
		response := helper.JsonResponse("Gagal menyimpan movie", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := movie.FormatMovie(newMovie)
	response := helper.JsonResponse("Penyimpanan movie berhasil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}

func (mb *MovieHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	movie, err := mb.movieService.DeleteMovie(id)

	if err != nil {
		response := helper.JsonResponse("Gagal mendapatkan movie", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("movie berhasil diambil", http.StatusOK, "success", movie)

	c.JSON(http.StatusOK, response)
}
