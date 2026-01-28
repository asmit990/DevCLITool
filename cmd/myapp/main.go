package main

import (
	"log"
	"os"

	"github.com/asmit990/myapp/internal/service"
	"github.com/asmit990/myapppkg/config"
	"github.com/asmit990/pkg/logger"
)

func main() {
	logr := logger.New()
	cfg := config.Load()

	logr.Println("app starting")

	svc := service.New(cfg, logr)

	if err := svc.Run(); err != nil {
		logr.Println("error:", err)
		os.Exit(1)
	}
}
