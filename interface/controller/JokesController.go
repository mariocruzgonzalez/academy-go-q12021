package controller

import (
	"academy-go-q12021/domain/model"
	"academy-go-q12021/usecase/repository"
	"github.com/labstack/echo"
	"net/http"

)


func GetRandomJoke(c echo.Context) error {
	var u model.ChuckJoke

	err, u := repository.GetRandomJoke(u)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}

	return c.JSON(http.StatusOK, u)
}

func GetAllCSVJokes(c echo.Context) error {
	u := map[int]model.ChuckJoke{}

	u, err := repository.GetAllCSVJokes(u)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}

	return c.JSON(http.StatusOK, u)
}
