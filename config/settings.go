package configuration

import (
	"log"
	"os"
	v "pinares/config/variables"
	"strings"
)

func LoadConfig() v.Config {
	env := os.Getenv("ENV")
	os.Setenv("CGO_ENABLED", "1")
	if env == "" {
		log.Fatal("ENV environment variable was not found")
	}
	jwtIssuer := os.Getenv("JWT_ISSUER")
	if jwtIssuer == "" {
		log.Fatal("JWT_ISSUER environment variable was not found")
	}
	globalKafkaHostPlain := os.Getenv("GLOBAL_KAFKA_HOST_PLAIN")

	config := v.Dev

	switch env {
	case "local":
		config = v.Local
	case "prod":
		config = v.Prod
	}

	if config.KafkaBrokers == nil {
		brokersList := strings.Split(globalKafkaHostPlain, ",")
		if brokersList[0] == "" {
			log.Fatal("GLOBAL_KAFKA_HOST_PLAIN environment variable was not found")
		}
		config.KafkaBrokers = brokersList
	}

	config.JwtIssuer = jwtIssuer

	return config
}

var Config_ = LoadConfig()
