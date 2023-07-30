package config

import "github.com/aws/aws-sdk-go/aws"

type AwsConfigurator struct {
	Region   string
	Endpoint string
	Config   *aws.Config
}

type RedisConfigurator struct {
	Host string
	Port string
	Db   string
	Pass string
}

type Global struct {
	// Global
	Mode  string
	Redis *RedisConfigurator
	Aws   *AwsConfigurator
}
