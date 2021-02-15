package cfg

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBConnectionString string `mapstructure:"DB_CONNECTION_STRING"`
	SecretKey          string `mapstructure:"SECRET_KEY"`
	S3BucketName       string `mapstructure:"S3_BUCKET_NAME"`
	S3Region           string `mapstructure:"S3_REGION"`
	S3APIKey           string `mapstructure:"S3_API_KEY"`
	S3Secret           string `mapstructure:"S3_SECRET"`
	S3Domain           string `mapstructure:"S3_DOMAIN"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
