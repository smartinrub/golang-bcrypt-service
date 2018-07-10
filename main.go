package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// PASSWORD hashed with BCrypt
type PASSWORD struct {
	HASH string `json:"password" binding:"required"`
}

func main() {
	router := SetupRouter()
	router.Run(":8089")
}

// SetupRouter creates endpoint for bcrypt
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/bcrypt", func(c *gin.Context) {
		var password PASSWORD
		c.BindJSON(&password)
		hash := hashPassword(password)
		fmt.Printf("Hashed value: %v\n", string(hash))
		c.String(201, string(hash))
	})

	return r
}

func hashPassword(password PASSWORD) string {
	fmt.Printf("Value: %v\n", password)
	rounds, _ := strconv.Atoi(os.Getenv("ROUNDS"))
	hash, _ := bcrypt.GenerateFromPassword([]byte(password.HASH), rounds)
	return string(hash)
}
