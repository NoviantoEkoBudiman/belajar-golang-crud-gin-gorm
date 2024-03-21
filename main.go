package main

import (
	"coba1/models"

	"coba1/controllers/mahasiswacontroller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.DbConnect()

	r.GET("/api/mahasiswa", mahasiswacontroller.Index)
	r.GET("/api/mahasiswa/:id", mahasiswacontroller.Show)
	r.POST("/api/mahasiswa", mahasiswacontroller.Create)
	r.PUT("/api/mahasiswa/:id", mahasiswacontroller.Update)
	r.DELETE("/api/mahasiswa/:id", mahasiswacontroller.Delete)
	r.Run()
}
