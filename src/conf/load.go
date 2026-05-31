package conf

import (
	"fmt"
	"log"
	"sync"
	"time"
	"todo_list/src/consts"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

type Config struct {
	NAME        string        `validate:"required"`
	HOST        string        `validate:"required"`
	PORT        int           `validate:"required,gt=0,lt=65536"`
	SECRET      string        `validate:"required,min=32"`
	GO_ENV      consts.GO_ENV `validate:"required,oneof=development production"`
	CORS_ORIGIN string        `validate:"required"`
	VERSION     string        `validate:"required"`
	// Rate limiter
	RATE_LIMIT_WINDOWS time.Duration `validate:"required,gt=0"`
	RATE_LIMIT_MAX_REQ int           `validate:"required,gt=0"`
	// Database
	DB_URL string `validate:"required"`
	// Jwt config
	JWT_ACCESS_EXP  time.Duration `validate:"required,gt=0"`
	JWT_REFRESH_EXP time.Duration `validate:"required,gt=0"`
	// Body parser limit
	BODY_LIMIT int       `validate:"required,gt=0"`
	IS_DEV     bool      `validate:"boolean"`
	START_TIME time.Time `validate:"required"`
	// Trust Proxy
	TRUST_PROXY  bool   `validate:"boolean"`
	PROXY_HEADER string `validate:"omitempty"`
}

var (
	Env  *Config
	once sync.Once
)

func init() {
	once.Do(func() {
		if err := godotenv.Load(".env"); err != nil {
			log.Printf("Error: .env file not found: %v\n", err)
		}
		GO_ENV := getEnv("GO_ENV", consts.DEVELOPMENT)
		Env = &Config{
			NAME:               getEnv("NAME"),
			PORT:               getEnvInt("PORT", 8000),
			HOST:               getEnv("HOST", "0.0.0.0"),
			VERSION:            "v1.0.0",
			GO_ENV:             GO_ENV,
			CORS_ORIGIN:        getEnv("CORS_ALLOW_ORIGIN"),
			SECRET:             getEnv("SECRET"),
			DB_URL:             getEnv("DB_URL"),
			START_TIME:         time.Now(),
			IS_DEV:             GO_ENV == consts.DEVELOPMENT,
			BODY_LIMIT:         int(getEnvByte("BODY_LIMIT")),
			JWT_ACCESS_EXP:     getEnvTime("JWT_ACCESS_EXP", "5m"),
			JWT_REFRESH_EXP:    getEnvTime("JWT_REFRESH_EXP", "24d"),
			RATE_LIMIT_MAX_REQ: getEnvInt("RATE_LIMIT_MAX_REQ", 100),
			RATE_LIMIT_WINDOWS: getEnvTime("RATE_LIMIT_WINDOWS", "15m"),
			TRUST_PROXY:        getEnvBool("TRUST_PROXY", false),
			PROXY_HEADER:       getEnv("PROXY_HEADER", fiber.HeaderXForwardedFor),
		}
		if err := validator.New().Struct(Env); err != nil {
			str := fmt.Sprintf("❌ Config validation failed: %v", err)
			panic(str)
		}
	})
}
