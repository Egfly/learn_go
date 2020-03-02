package config

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
)

var (
	AppConfig *Config
)

type Config struct {
	LogLevel       string
	LogPath        string
	CollectsConfig []CollectConfig
	ChannelSize    int
	KafkaAddr      string
}

type CollectConfig struct {
	LogPath string
	Topic   string
}

func LoadConfig(configType, path string) (err error) {
	conf, err := config.NewConfig(configType, path)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	AppConfig = &Config{}
	AppConfig.LogLevel = conf.String("logs::log_level")
	if len(AppConfig.LogLevel) == 0 {
		AppConfig.LogLevel = "debug"
	}

	AppConfig.LogPath = conf.String("logs::log_path")
	if len(AppConfig.LogPath) == 0 {
		AppConfig.LogPath = "logs"
	}

	AppConfig.ChannelSize, err = conf.Int("logs::channel_size")
	if err != nil {
		AppConfig.ChannelSize = 100
	}

	AppConfig.KafkaAddr = conf.String("kafka::server_addr")
	if len(AppConfig.KafkaAddr) == 0 {
		AppConfig.KafkaAddr = "192.168.31.239:9092"
	}

	err = loadCollectConfig(conf)
	if err != nil {
		fmt.Println("load collect config failed, err:", err)
		return
	}
	return
}

func loadCollectConfig(conf config.Configer) (err error) {
	var cc CollectConfig
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect log_path")
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid collect topic")
	}
	AppConfig.CollectsConfig = append(AppConfig.CollectsConfig, cc)
	return
}
