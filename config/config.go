/*

 */

package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config holds the config data object with the actual data. Methods written over Config object give
// access to actual data.
// This is to make configurations read-only and also allow easy parsing of configuration files.
type Config struct {
	config configData
}

type configData struct {
	// http
	HTTPPort int    `mapstructure:"http_port"`
	GinMode  string `mapstructure:"gin_mode"`

	// log
	LogLevel        logrus.Level `mapstructure:"log_level"`
	LogFileLocation string       `mapstructure:"log_file_location"`

	// database
	DBType           string `mapstructure:"db_type"`
	DBConnectionPath string `mapstructure:"db_path"`
}

func NewConfig() (*Config, error) {

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")

	// where to look for
	viper.AddConfigPath("/etc/nftactions/api/") // production config path
	viper.AddConfigPath("./config")             // dev config path
	viper.AddConfigPath("../config")            // dev config path

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, err
	}

	// Override config file based on ENV variables e.g. DB_TYPE=postgres
	viper.AutomaticEnv()

	config := configData{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &Config{config: config}, nil
}

func (c *Config) HTTPPort() int {
	return c.config.HTTPPort
}

func (c *Config) GinMode() string {
	return c.config.GinMode
}

func (c *Config) LogFileLocation() string {
	return c.config.LogFileLocation
}

func (c *Config) LogLevel() logrus.Level {
	return c.config.LogLevel
}

func (c *Config) DBCredentials() []string {
	return []string{
		c.config.DBType,
		c.config.DBConnectionPath,
	}
}
