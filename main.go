package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)
func main(){
	e := echo.New()
	e.GET("/",func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Halo selamat datang Ady, selamat belajar",
		})
	})
	e.Start(":8000")
	fmt.Println("server berjalan di port 8000")
}