package config

import "os"

type Configuration struct {

}

func (Configuration) GetPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	return "8080"
}
