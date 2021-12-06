package router

import (
	"github.com/abhishek-singh/handler"
	"github.com/gin-gonic/gin"
)

func Routes(app *gin.Engine) {
	/* router secton here */
	app.GET("/emp", handler.GetAllEmpoyee)
	app.POST("/emp", handler.CreateProduct)
	app.GET("/emp/:id", handler.FindOneProduct)
	app.DELETE("/emp/:id", handler.DeleteOneProduct)
	app.PUT("/emp/:id", handler.UpdataeProduct)
}
