package main

import (
	"log"
	"path/filepath"

	"github.com/velvetriddles/bank-app/config"
	"github.com/velvetriddles/bank-app/internal/app"
	"github.com/velvetriddles/bank-app/internal/di"
)

func main() {
	cmdDir, err := filepath.Abs(filepath.Dir(""))
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	rootDir := filepath.Dir(cmdDir)

	cfg, err := config.LoadConfig(rootDir)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	container := di.NewContainer(cfg)
	if err := container.Init(); err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	application := app.NewApp(container)
	if err := application.Run(); err != nil {
		log.Fatalf("Failed to run app: %v", err)
	}
}
