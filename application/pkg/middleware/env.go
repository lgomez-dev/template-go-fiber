package middleware

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	envErr := godotenv.Load("app.env")
	if envErr != nil {
		fmt.Println("Error loading env file")
	}
}
