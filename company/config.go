package company

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config map somes variables defaults from application
type Config struct {
	APP  string `envconfig:"APP_NAME" default:"compnay"`
	Port int    `envconfig:"APP_PORT" default:"8083"`

	AWSRegion      string `envconfig:"COMPANY_AWS_REGION" default:"us-east-1"`
	AWSMetadataURL string `envconfig:"COMPANY_AWS_METADATA_URL" default:"http://169.254.169.254:80/latest"`

	DynamodbEndpoint string `envconfig:"COMPANY_DYNAMODB_ENDPOINT" default:"http://localhost:4569"`
}

//NewConfig config constructor
func NewConfig() *Config {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
