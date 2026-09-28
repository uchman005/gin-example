package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func get(c *gin.Context) {
	Success(c, map[string]string{
		"success": "true",
	})
}
func me(c *gin.Context) {
	Success(c, "Uche")
}

func pathParam(c *gin.Context) {
	name := c.Param("name")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	action := c.Param("action")
	if action != "" {
		message := fmt.Sprintf("Hello %s action is %s", name, action)
		Success(c, message)
		return
	}
	c.String(http.StatusOK, "Hello %s %d", name, limit)
}

func welcome(c *gin.Context) {
	firstName := c.DefaultQuery("first", "Uche")
	lastName := c.Query("last")

	c.String(http.StatusOK, "Welcome %s %s", firstName, lastName)
}

func final(c *gin.Context) {
	c.String(http.StatusAccepted, "Redirected successfully")
}
