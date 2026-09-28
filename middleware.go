package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	name string
	age  int
}

func (user User) toString() {
	fmt.Printf("%s with age %d", user.name, user.age)
}
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("MiddleWare Active")
		user := User{
			name: "Uche",
			age:  30,
		}
		ctx.Set("name", user.name)
		ctx.Set("age", user.age)
		ctx.Next()
	}
}

type AppError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	ErrNotFound = &AppError{
		Status:  http.StatusNotFound,
		Code:    "NOT_FOUND",
		Message: "resource not found"}
	ErrUnauthorized = &AppError{
		Status:  http.StatusUnauthorized,
		Code:    "UNAUTHORIZED",
		Message: "authentication required"}
	ErrBadRequest = &AppError{
		Status:  http.StatusBadRequest,
		Code:    "BAD_REQUEST",
		Message: "invalid request"}
	ErrorTooManyRequest = &AppError{
		Status:  http.StatusTooManyRequests,
		Code:    "TOO_MANY_REQUEST",
		Message: "Too many request to the server",
	}
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		var appErr *AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Status, gin.H{
				"success": false,
				"error":   gin.H{"code": appErr.Code, "message": appErr.Message},
				"status":  appErr.Status,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   gin.H{"code": "INTERNAL", "message": "an unexpected error occurred"},
				"status":  http.StatusInternalServerError,
			})
		}
	}
}
