package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
}

var configuration Config = Config{}

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

	configuration = Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
	}

}

func GetConfig() Config {
	loadConfig()
	return configuration
}
