package mahasiswacontroller

import (
	"coba1/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var mahasiswa []models.Mahasiswa
	models.DB.Find(&mahasiswa)
	c.JSON(http.StatusOK, gin.H{"data": mahasiswa})
}

func Show(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")
	if err := models.DB.First(&mahasiswa, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Data not found",
			})
			return
		default:
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{"message": err.Error()},
			)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data": mahasiswa,
	})
}

func Create(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&mahasiswa)
	c.JSON(http.StatusOK, gin.H{
		"message": "Data Berhasil Disimpan",
		"data":    mahasiswa,
	})

}

func Update(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.Model(&mahasiswa).Where("mahasiswa_id = ?", id).Updates(&mahasiswa).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate"})
}

func Delete(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")

	if err := models.DB.Delete(&mahasiswa, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
