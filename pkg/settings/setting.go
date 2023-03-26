package settings

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Server struct {
	RunMode      string        `mapstructure:"run_mode"`
	HTTPPort     int           `mapstructure:"http_port"`
	ReadTimeOut  time.Duration `mapstructure:"read_timeout"`
	WriteTimeOut time.Duration `mapstructure:"write_timeout"`
}

type App struct {
	JwtSecret         string   `mapstructure:"jwt_secret"`
	PageSize          int      `mapstructure:"page_size"`
	PrefixURL         string   `mapstructure:"prefix_url"`
	RuntimeRootPath   string   `mapstructure:"runtime_root_path"`
	ImageSavePath     string   `mapstructure:"image_save_path"`
	ImageMaxSize      int      `mapstructure:"image_size"`
	ImageAllowExts    []string `mapstructure:"image_allow_ext"`
	ExportSavePath    string   `mapstructure:"export_save_path"`
	QrCodeSavePath    string   `mapstructure:"qr_code"`
	LogSavePath       string   `mapstructure:"log_save_path"`
	LogSaveName       string   `mapstructure:"log_save_name"`
	LogFileExt        string   `mapstructure:"log_file_ext"`
	TimeFormat        string   `mapstructure:"time_format"`
	Issuer            string   `mapstructure:"issuer"`
	UrlForgotPassword string   `mapstructure:"url_forgot_password"`
	UrlVerityUser     string   `mapstructure:"url_verity_user"`
}

type Database struct {
	Type        string `mapstructure:"type"`
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Name        string `mapstructure:"name"`
	TablePrefix string `mapstructure:"table_prefix"`
}

type RedisDB struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	Key      string `mapstructure:"key"`
	Password string `mapstructure:"password"`
}

type FileConfig struct {
	JWTExpired int       `mapstructure:"expire_jwt"`
	Server     *Server   `mapstructure:"server"`
	App        *App      `mapstructure:"app"`
	Database   *Database `mapstructure:"database"`
	RedisDB    *RedisDB  `mapstructure:"redisdb"`
}

var AppConfigSetting = &FileConfig{}

func Setup(configFile string) {
	now := time.Now()
	viper.SetConfigFile(configFile)
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
