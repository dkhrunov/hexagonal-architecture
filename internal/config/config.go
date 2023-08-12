package config

import (
	"os"
)

type PortServiceConfig struct {
	HTTPAddr string
}

type Config struct {
	PortService PortServiceConfig
}

func Read() Config {
	return Config{
		PortService: PortServiceConfig{
			HTTPAddr: readEnv("PORT_SERVICE_HTTP_ADDR", ""),
		},
	}
}

// Helper function to read an environment or return a fallback value
func readEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Uncomment this if necessary -->
// // Helper function to read an environment variable into integer or return a fallback value
// func readEnvAsInt(key string, fallback int) int {
// 	envValue := readEnv(key, "")
// 	if value, err := strconv.Atoi(envValue); err == nil {
// 		return value
// 	}
// 	return fallback
// }

// Uncomment this if necessary -->
// // Helper function to read an environment variable into boolean or return a fallback value
// func readEnvAsBool(key string, fallback bool) bool {
// 	envValue := readEnv(key, "")
// 	if value, err := strconv.ParseBool(envValue); err == nil {
// 		return value
// 	}
// 	return fallback
// }

// Uncomment this if necessary -->
// // Helper to read an environment variable into a string slice or return fallback value
// func readEnvAsSlice(key string, fallback []string, sep string) []string {
// 	envValue := readEnv(key, "")
// 	if envValue == "" {
// 		return fallback
// 	}
// 	return strings.Split(envValue, sep)
// }
