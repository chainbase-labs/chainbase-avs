package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/chainbase-avs/cli/cmd"
)

func main() {

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	logger := slog.New(handler)

	// set default logger
	slog.SetDefault(logger)

	// err := godotenv.Load()
	// if err != nil {
	// 	slog.Error("failed to load .env file", "error", err)
	// }
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
