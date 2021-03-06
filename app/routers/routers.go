package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/idriscahyono/bookstore/app/controllers"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

func InitRouter() {
	port := viper.Get("PORT").(int)
	r := gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic("Cannot start service" + err.Error())
	}
}
