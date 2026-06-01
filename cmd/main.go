package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo_list/src/conf"
	"todo_list/src/db"
	"todo_list/src/lib/utils"
	_ "todo_list/src/logger"
	"todo_list/src/server"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := server.New()
	uri := fmt.Sprintf("%s:%d", conf.Env.HOST, conf.Env.PORT)
	utils.PrintBanner(uri, conf.Env.GO_ENV)

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(uri, fiber.ListenConfig{
			EnablePrefork:         false,
			DisableStartupMessage: true,
		}); err != nil && errors.Is(err, net.ErrClosed) {
			log.Panic(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(done)

	<-done // This blocks the main thread until an interrupt is received
	fmt.Println("\nGracefully shutting down...")
	if err := app.ShutdownWithTimeout(10 * time.Second); err != nil {
		fmt.Printf("Server shutdown failed: %v\n", err)
	}

	// Your cleanup tasks go here
	fmt.Println("Running cleanup tasks...")
	db.Pool.Close()
	fmt.Println("Cleanup completed. Bye 👋")
}
