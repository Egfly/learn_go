package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	config2 "oldBoyGo/src/day11/logagent/config"
)

func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "alert":
		return logs.LevelAlert
	case "info":
		return logs.LevelInfo
	case "error":
		return logs.LevelError
	default:
		return logs.LevelDebug
	}
}

func initLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = config2.AppConfig.LogPath
	config["level"] = convertLogLevel(config2.AppConfig.LogLevel)
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("init logger failed, err:", err)
		return
	}

	err = logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}
