package main

import (
	"flag"
	"log"

	"ps_info/pkg/config"
	"ps_info/pkg/logger"
)

func main() {
	configFilePath := flag.String("config", "", "config file ABS path")
	flag.Parse()

	// load project config
	projectConfigContent, err := config.InitConfig(*configFilePath)
	if err != nil {
		log.Println("[Init Config] Failed to init project config, err: ", err.Error())
		return
	}
	log.Println("[Init Config] Succeed to init project config......")

	// init project logger
	logConfigContent := projectConfigContent.LogConfig
	err = logger.InitLog(logConfigContent)
	if err != nil {
		log.Println("[Init Logger] Failed to init logger, err: ", err.Error())
		return
	}
	log.Println("[Init Logger] Succeed to init logger......")
	// pkg.Output()
}
