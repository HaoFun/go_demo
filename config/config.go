package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/lexkong/log"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	if err := c.initConfig(); err != nil {
		return err
	}

	c.initLog()

	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("dev")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("go_demo")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:       viper.GetString("log.writers"),
		LoggerLevel:   viper.GetString("log.logger_level"),
		LoggerFile:    viper.GetString("log.logger_file"),
		LogFormatText: viper.GetBool("log.log_format_text"),
		RollingPolicy: viper.GetString("log.rollingPolicy"),
		LogRotateDate: viper.GetInt("log.log_rotate_date"),
		LogRotateSize: viper.GetInt("log.log_rotate_size"),
		LogBackupCount:viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}

// 監聽配置文件變化並熱加載
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}