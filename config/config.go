package config

import (
	"os"
	"regexp"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	Jaeger   JaegerConfig   `mapstructure:"jaeger"`
}

type ServerConfig struct {
	Host           string `mapstructure:"host"`
	Port           string `mapstructure:"port"`
	Debug          bool   `mapstructure:"debug"`
	ReadTimeout    int    `mapstructure:"readTimeout"`
	WriteTimeout   int    `mapstructure:"writeTimeout"`
	ContextTimeout int    `mapstructure:"contextTimeout"`
	Timezone       string `mapstructure:"timezone"`
	Location       *time.Location
}

type DatabaseConfig struct {
	Adapter         string `mapstructure:"adapter"`
	Host            string `mapstructure:"host"`
	Username        string `mapstructure:"username"`
	Db              string `mapstructure:"db"`
	Password        string `mapstructure:"password"`
	Port            int    `mapstructure:"port"`
	MaxConns        int    `mapstructure:"maxConns"`
	MaxLiftimeConns int    `mapstructure:"maxLiftimeConns"`
	SSL             bool   `mapstructure:"ssl"`
}

type LoggerConfig struct {
	Development       bool   `mapstructure:"development"`
	DisableCaller     bool   `mapstructure:"disableCaller"`
	DisableStacktrace bool   `mapstructure:"disableStacktrace"`
	Encoding          string `mapstructure:"encoding"`
	Level             string `mapstructure:"level"`
	Filename          string `mapstructure:"filename"`
	FileMaxSize       int    `mapstructure:"fileMaxSize"`
	FileMaxAge        int    `mapstructure:"fileMaxAge"`
	FileMaxBackups    int    `mapstructure:"fileMaxBackups"`
	FileIsCompress    bool   `mapstructure:"fileIsCompress"`
}

type JaegerConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	ServiceName string `mapstructure:"serviceName"`
	LogSpans    bool   `mapstructure:"logSpans"`
}

func GetConf() (*Config, error) {
	re := regexp.MustCompile(`\$\{([^{}]+)\}`)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if re.Match([]byte(value)) {
			env := string(re.ReplaceAll([]byte(value), []byte("$1")))
			viper.Set(k, os.Getenv(env))
		}

	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	loc, err := time.LoadLocation(config.Server.Timezone)
	if err != nil {
		return nil, err
	}
	config.Server.Location = loc

	return &config, nil
}
