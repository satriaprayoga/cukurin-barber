package settings

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Server struct {
	RunMode      string        `mapstucture:"run_mode"`
	HTTPPort     int           `mapstructure:"http_port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

type FileConfig struct {
	Server *Server `mapstructure:"server"`
}

var AppConfigSetting = &FileConfig{}

func Setup() {
	now := time.Now()
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fail to parse 'config.json': %v", err)
	}
	err = viper.Unmarshal(AppConfigSetting)
	if err != nil {
		log.Fatalf("Fail to Unmarshall 'config.json': %v", err)
	}
	timeSpent := time.Since(now)
	log.Printf("Config setting is ready in %v", timeSpent)
}
