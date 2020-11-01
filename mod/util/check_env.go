package util

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func CheckEnv(e *echo.Echo) {
	if err := godotenv.Load(); err != nil {
		e.Logger.Fatalf("Error godotenv :%v", err)
	}
	if os.Getenv("BASKETBALL_KEY") == "" {
		e.Logger.Fatal("failed load env BASKETBALL_KEY")
	}
	if os.Getenv("SOCCER_KEY") == "" {
		e.Logger.Fatal("failed load env SOCCER_KEY")
	}
}
