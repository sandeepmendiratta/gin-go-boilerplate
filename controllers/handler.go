
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandeepmendiratta/gin-go-boilerplate/models"
)
func HandleGet(c *gin.Context) {
message, _ := c.GetQuery("m")
c.String(http.StatusOK, "Get works! you sent: "+message)
//localhost:8080/api/?m=sandeep
}
func HandleVerification(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		// setup headers
		c.Header("Allow", "POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "origin, content-type, accept")
		c.Header("Content-Type", "application/json")
		c.Status(http.StatusOK)
	} else if c.Request.Method == "POST" {
		var u models.User
		c.BindJSON(&u)
		c.JSON(http.StatusOK, gin.H{
			"user": u.Username,
			"pass": u.Password,
		})
	}
}