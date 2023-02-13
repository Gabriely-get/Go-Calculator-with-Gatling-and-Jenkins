package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

var history = []Operation{}

func main() {
	createLogDirectory()
	infoLogger("Application started. Running at port 8000")

	router := gin.Default()
	router.GET("/calc/sum/:num1/:num2", handlerSum)
	router.GET("/calc/sub/:num1/:num2", handlerSub)
	router.GET("/calc/mult/:num1/:num2", handlerMult)
	router.GET("/calc/div/:num1/:num2", handlerDiv)
	router.GET("/calc/history", handlerHistory)

	err := router.Run(":8000")
	if err != nil {
		fatalLogger(fmt.Sprintf("Error on starting application: err: %v", err))
	} 

	
}
