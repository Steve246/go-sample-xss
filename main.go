package main

import (
	"fmt"
	"golang-sample-xss/model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//CROSS SITE SCRIPTING

func main() {

	users := make([]model.User, 0)

	apiHost := os.Getenv("API_HOST")  //localhost
	apiPort := os.Getenv("API_PORT") // 8888

	listenAddress := fmt.Sprintf("%s:%s", apiHost, apiPort)

	routerEngine := gin.Default()

	routerGroup := routerEngine.Group("/api")


	routerGroup.POST("/user", func(ctx *gin.Context){

		var newUser model.User

		if err := ctx.ShouldBindJSON(&newUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H {
				"error": err.Error(),
			})
			return 
		}

		users = append(users, newUser)

		ctx.JSON(http.StatusOK, gin.H {
			"message": "Success",
			"data": newUser,
		})

	})

	err := routerEngine.Run(listenAddress)

	if err != nil {
		panic(err)
	}



}