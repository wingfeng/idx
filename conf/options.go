package conf

import (
	"log/slog"

	"github.com/spf13/viper"
)

type Option struct {
	LogLevel       string
	Port           int
	Driver         string
	Connection     string
	PrivateKeyPath string
	PublicKeyPath  string
	SyncDB         bool
	HTTPScheme     string //服务器有可能放在nginx做了SSL卸载，所以无法直接判断是https还是http
	RedisHost      string
	RedisPort      int
	RedisDB        int
}

var Options *Option

func init() {
	Options = initConfig()
}

func initConfig() *Option {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../conf/")
	viper.AddConfigPath("./conf/")

	viper.AllowEmptyEnv(true)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file

		slog.Error("读取配置文件错误,系统将尝试从Env中读取配置:", "error", err)
	}
	viper.SetEnvPrefix("IDX")
	viper.AutomaticEnv()

	opts := &Option{}

	err = viper.Unmarshal(opts)
	if err != nil {
		slog.Error("读取配置错误:", "error", err)
	}

	return opts
}
