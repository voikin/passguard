package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/voikin/passguard/internal/controllers"
	"github.com/voikin/passguard/internal/services/password"
	"github.com/voikin/passguard/internal/usecases"
)

func main() {
	commonPatternsFilename := os.Getenv("PATTERNS_FILE")
	if commonPatternsFilename == "" {
		log.Fatal("empty patterns file path")
	}

	httpPort := os.Getenv("HTTP_PORT")
	if commonPatternsFilename == "" {
		log.Fatal("empty http port")
	}

	pwdService, err := password.New(commonPatternsFilename)
	if err != nil {
		log.Fatalf("create password service: %v", err)
	}

	useCases := usecases.New(pwdService)
	controllers := controllers.New(useCases)

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/evaluate", controllers.EvaluatePasswordHandler)
	}

	err = r.Run(fmt.Sprintf(":%s", httpPort))
	if err != nil {
		log.Fatalf("start http server: %v", err)
	}
}
