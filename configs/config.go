package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST    		string `mapstructure:"DB_HOST"`
	DB_NAME  		string `mapstructure:"DB_NAME"`
	DB_USER  		string `mapstructure:"DB_USER"`
	DB_PASSWORD 	string `mapstructure:"DB_PASSWORD"`
	DB_PORT 		string `mapstructure:"DB_PORT"`
	JWT_SECRET		string `mapstructure:"JWT_SECRET"`
	SUPABASE_URL    string `mapstructure:"SUPABASE_URL"`
	SUPABASE_API_KEY string `mapstructure:"SUPABASE_API_KEY"`
	SUPABASE_BUCKET string `mapstructure:"SUPABASE_BUCKET"`
	
	// SMTP settings
	SMTP_HOST       string `mapstructure:"SMTP_HOST"`
	SMTP_PORT       int    `mapstructure:"SMTP_PORT"`
	EMAIL_FROM      string `mapstructure:"SMTP_USER"`
	EMAIL_PASSWORD  string `mapstructure:"SMTP_PASSWORD"`
}

func NewConfig() *Config {
	config := &Config{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("❌ Error reading config file", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		log.Fatalln("❌ Unable to decode into struct", err)
	}

	// Add default values for SMTP port if not set in env
	if config.SMTP_PORT == 0 {
		config.SMTP_PORT = 587 // Default SMTP port
	}

	return config
}
