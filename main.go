package main

import (
	"log"
	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/config"
	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/handlers"
	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/middleware"
	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/repository"
	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()
}