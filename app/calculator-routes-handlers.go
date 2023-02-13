package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
	WL string = "w"
	IL string = "i"
	EL string = "e"
	FL string = "f"
)

func handlerSum(c *gin.Context) {
	floatNum1, err := convertToFloat( c.Param("num1") )
	if err != nil {
		returnValueConvertingBadRequest(c, c.Param("num1"), err)
		return
	}

	floatNum2, err := convertToFloat( c.Param("num2") )
	if err != nil {
		returnValueConvertingBadRequest(c, c.Param("num2"), err)
		return
	}

	sumResult := sum(floatNum1, floatNum2)

	newOperation := Operation {
		Time: time.Now(),
		Operation: "ADDITION",
		Result: returnResult(floatNum1, floatNum2, sumResult),
 	}
	addToHistory(newOperation)

	c.IndentedJSON(http.StatusOK, returnSuccess(sumResult))
	loggingRequest(c, 200, IL, " ")
}

func handlerSub(c *gin.Context) {
	floatNum1, err := convertToFloat( c.Param("num1") )
	if err != nil {
		returnValueConvertingBadRequest(c, c.Param("num1"), err)
		return
	}

	floatNum2, err := convertToFloat( c.Param("num2") )
	if err != nil {
		returnValueConvertingBadRequest(c, c.Param("num2"), err)
		return
	}

	subResult := sub(floatNum1, floatNum2)

	newOperation := Operation {
		Time: time.Now(),
		Operation: "SUBTRACTION",
		Result: returnResult(floatNum1, floatNum2, subResult),
 	}
	addToHistory(newOperation)

	c.IndentedJSON(http.StatusOK, returnSuccess(subResult))
	loggingRequest(c, 200, IL, " ")
}

func handlerMult(c *gin.Context) {
	floatNum1, err := convertToFloat( c.Param("num1") )
	if err != nil {
		returnValueConvertingBadRequest(c, c.Param("num1"), err)
		return
	}

	floatNum2, err := convertToFloat( c.Param("num2") )
	if err != nil {
		returnValueConvertingBadRequest(c, c.Param("num2"), err)
		return
	}

	multResult := mult(floatNum1, floatNum2)

	newOperation := Operation {
		Time: time.Now(),
		Operation: "MULTIPLICATION",
		Result: returnResult(floatNum1, floatNum2, multResult),
 	}
	addToHistory(newOperation)

	c.IndentedJSON(http.StatusOK, returnSuccess(multResult))
	loggingRequest(c, 200, IL, " ")
}

func handlerDiv(c *gin.Context) {
	floatNum1, err := convertToFloat( c.Param("num1") )
	if err != nil {
		returnValueConvertingBadRequest(c, c.Param("num1"), err)
		return
	}

	floatNum2, err := convertToFloat( c.Param("num2") )
	if err != nil {
		returnValueConvertingBadRequest(c, c.Param("num2"), err)
		return
	}

	if floatNum2 == 0 {
		message := "Mathematical indeterminacy: division by 0"

		c.IndentedJSON(http.StatusBadRequest, returnError(message))
		loggingRequest(c, 400, WL, message)
		return
	}

	divResult := div(floatNum1, floatNum2)

	newOperation := Operation {
		Time: time.Now(),
		Operation: "DIVISION",
		Result: returnResult(floatNum1, floatNum2, divResult),
 	}
	addToHistory(newOperation)

	c.IndentedJSON(http.StatusOK, returnSuccess(divResult))
	loggingRequest(c, 200, IL, " ")
}

func handlerHistory(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, history)
	loggingRequest(c, 200, IL, " ")
}

func returnResult(param1 float64, param2 float64, result float64) string {
	return fmt.Sprintf("%.2f / %.2f = %.2f", param1, param2, result) 
}

func returnValueConvertingBadRequest(c *gin.Context, value string, err error) {
	message := fmt.Sprintf("Could not convert value %s", value)

	loggingRequest(c, 400, WL, fmt.Sprintf(message+": %s", err))
	c.IndentedJSON(http.StatusBadRequest, returnError(message))
}

func loggingRequest(c *gin.Context, status int, logger string, message string) {

	message = fmt.Sprintf("%s %s %s %s", c.Request.Method, fmt.Sprint(status), c.Request.Host+c.Request.URL.Path, message)

	if (logger == IL) {
		infoLogger(message)
	}

	if (logger == WL) {
		warningLogger(message)
	}

	if (logger == EL) {
		errorLogger(message)
	}

	if (logger == FL) {
		fatalLogger(message)
	}

}