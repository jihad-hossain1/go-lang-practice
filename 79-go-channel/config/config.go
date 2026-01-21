package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	SslMode  string
}

var dbConfiguration *Config

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load then env variables", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	fmt.Println(version)
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http port are required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("port must be number")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("Jwt secret key required")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("database host is required")
		os.Exit(1)
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("database port is required")
		os.Exit(1)
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("database name is required")
		os.Exit(1)
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("database USER is required")
		os.Exit(1)
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("database password is required")
		os.Exit(1)
	}
	dbSslMode := os.Getenv("DB_SSL_MODE")
	if dbSslMode == "" {
		fmt.Println("database database ssl mode is required")
		os.Exit(1)
	}

	dPort, err := strconv.Atoi(dbPort)
	// convertSslMode, err := strconv.ParseBool(dbSslMode)
	if err != nil {
		fmt.Println("Invalid ssl mode")
		os.Exit(1)
	}

	dbConfiguration := &DBConfig{
		Host:     dbHost,
		Port:     dPort,
		SslMode:  dbSslMode,
		Password: dbPassword,
		Name:     dbName,
		User:     dbUser,
	}

	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfiguration,
	}

}

func GetConfig() *Config {
	if configuration == nil {
		// first time load config
		loadConfig()
	}
	// second time
	return configuration
}
