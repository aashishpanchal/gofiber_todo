package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo_list/boot/conf"
	"todo_list/cmd/server"
	"todo_list/logs"
	"todo_list/pkgs/utils"

	"github.com/gofiber/fiber/v3"
)

func main() {
	conf.Init()
	logs.Init()
	app := server.New()
	uri := fmt.Sprintf("%s:%d", conf.Env.HOST, conf.Env.PORT)
	// Print banner
	utils.PrintBanner(uri, conf.Env.GO_ENV)

	// Run server in goroutine
	go func() {
		if err := app.Listen(uri, fiber.ListenConfig{
			EnablePrefork:         false,
			DisableStartupMessage: true,
		}); err != nil {
			fmt.Printf("❌ Server failed to start: %v\n", err)
			os.Exit(1)
		}
	}()

	// Create channel to listen for OS signals
	done := make(chan os.Signal, 1)
	// Listen for SIGINT & SIGTERM (docker/podman compatible)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done // Block until signal received

	fmt.Println() // New line after ^C
	fmt.Println("🛑 Shutting down server...")

	// Create timeout context
	timer := time.AfterFunc(5*time.Second, func() {
		fmt.Println("⚠️ Shutdown timeout reached. Forcing exit.")
		os.Exit(1)
	})
	defer timer.Stop()

	// Gracefully shutdown
	if err := app.Shutdown(); err != nil {
		fmt.Printf("❌ Server shutdown failed: %v\n", err)
	} else {
		fmt.Println("✅ Server gracefully stopped")
	}

	// Add cleanup logic here
	fmt.Println("🧹 Cleanup completed. Bye 👋")
}
