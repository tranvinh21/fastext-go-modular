package config

import (
	"log"

	"github.com/spf13/viper"
)

var Envs = LoadConfig()

type Config struct {
	DBConfig     DBConfig        `mapstructure:"database"`
	ServerConfig ServerConfig    `mapstructure:"server"`
	AuthConfig   AuthConfig      `mapstructure:"auth"`
	Whitelist    WhitelistConfig `mapstructure:"whitelist"`
	S3Config     S3Config        `mapstructure:"s3_config"`
}

type DBConfig struct {
	URL string `mapstructure:"DATABASE_URL"`
}

type ServerConfig struct {
	Port string `mapstructure:"PORT"`
}

type AuthConfig struct {
	AccessTokenSecret  string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

type WhitelistConfig struct {
	Domains []string `mapstructure:"WHITELIST_DOMAINS"`
}

type S3Config struct {
	AccessKeyID     string `mapstructure:"S3_ACCESS_KEY_ID"`
	SecretAccessKey string `mapstructure:"S3_SECRET_ACCESS_KEY"`
	BucketName      string `mapstructure:"S3_BUCKET_NAME"`
}

func LoadConfig() *Config {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")    // Current directory
	viper.SetConfigName(".env") // Name of config file (without extension)
	viper.SetConfigType("env")  // Type of config file

	viper.SetDefault("PORT", "3000")
	viper.SetDefault("WHITELIST_DOMAINS", "*")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Error reading config file: %v\n", err)
	}

	requiredEnvs := []string{
		"DATABASE_URL",
		"ACCESS_TOKEN_SECRET",
		"REFRESH_TOKEN_SECRET",
		"S3_ACCESS_KEY_ID",
		"S3_SECRET_ACCESS_KEY",
		"S3_BUCKET_NAME",
	}

	for _, env := range requiredEnvs {
		if viper.GetString(env) == "" {
			log.Fatalf("Environment variable %s is not set", env)
		}
	}

	return &Config{
		DBConfig: DBConfig{
			URL: viper.GetString("DATABASE_URL"),
		},
		ServerConfig: ServerConfig{
			Port: viper.GetString("PORT"),
		},
		AuthConfig: AuthConfig{
			AccessTokenSecret:  viper.GetString("ACCESS_TOKEN_SECRET"),
			RefreshTokenSecret: viper.GetString("REFRESH_TOKEN_SECRET"),
		},
		Whitelist: WhitelistConfig{
			Domains: viper.GetStringSlice("WHITELIST_DOMAINS"),
		},
		S3Config: S3Config{
			AccessKeyID:     viper.GetString("S3_ACCESS_KEY_ID"),
			SecretAccessKey: viper.GetString("S3_SECRET_ACCESS_KEY"),
			BucketName:      viper.GetString("S3_BUCKET_NAME"),
		},
	}
}
