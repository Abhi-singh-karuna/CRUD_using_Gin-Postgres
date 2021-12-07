package handler

import (
	"net/http"

	"github.com/abhishek-singh/database"
	"github.com/abhishek-singh/model"
	"github.com/gin-gonic/gin"
)

func GetAllEmpoyee(c *gin.Context) {

	db := database.DB

	var data []model.Employee

	db.Find(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
	return
}
func CreateProduct(c *gin.Context) {

	db := database.DB

	var data = new(model.Employee)
	//fmt.Println(data.Names)
	c.BindJSON(&data)

	db.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})

	return
}

func FindOneProduct(c *gin.Context) {
	id := c.Param("id")

	db := database.DB

	var data model.Employee

	db.Find(&data, id)

	/*  */
	if data.Username == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "No product found with ID",
			"data":    nil,
		})
		return
	}
	/*  */

	c.JSON(http.StatusOK, gin.H{
		"message": "selected data are ",
		"data":    data,
	})

	return

}

func DeleteOneProduct(c *gin.Context) {
	id := c.Param("id")

	db := database.DB

	var data model.Employee

	db.Find(&data, id)

	db.Delete(&data)
	/*  */
	if data.Username == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "No product found with ID",
			"data":    nil,
		})
		return

	}
	/*  */

	c.JSON(http.StatusOK, gin.H{
		"message": "selected data are deleted ",
	})
	return
}

func UpdataeProduct(c *gin.Context) {
	var s = new(model.Employee)
	c.BindJSON(&s)

	id := c.Param("id")

	db := database.DB

	var data = new(model.Employee)

	db.Find(&data, id)

	data.Names = s.Names
	data.Username = s.Username
	data.Email = s.Email

	if s.Names == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "data not found witht this ID",
		})

	}

	db.Update(&data)
	db.Save(&data)

	c.JSON(http.StatusOK, gin.H{
		"message": "data updated  sucessfull",
		"data":    data,
	})

	return

}
