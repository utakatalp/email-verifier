package main

import (
	"github.com/gin-gonic/gin"
	"github.com/utakatalp/email-verifier/db"
	"github.com/utakatalp/email-verifier/handlers"
)

func main() {

	db.ConnectDB()

	r := gin.Default()
	r.GET("/activate", handlers.ActivateEmail)
	r.POST("/verify", handlers.VerifyEmail)

	r.Run(":8080")
}
