package main

import (
	"academy-go-q12021/infraesctructure/router"
	"fmt"
	"github.com/labstack/echo/middleware"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"


)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	router.NewRouter(e)

	fmt.Println("Server listen at http://localhost")
	if err := e.Start(":" ); err != nil {
		log.Fatalln(err)
	}
}
