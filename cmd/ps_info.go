package main

import (
	"flag"
	"log"
	"time"

	"ps_info/pkg/config"
	"ps_info/pkg/logger"
)

func main() {

	timeStart := time.Now().Nanosecond()

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
	serverConfigContent := projectConfigContent.ServerConfig
	err = logger.InitLog(logConfigContent, serverConfigContent)
	if err != nil {
		log.Println("[Init Logger] Failed to init logger, err: ", err.Error())
		return
	}
	log.Println("[Init Logger] Succeed to init logger......")

	// pkg.Output()

	timeEnd := time.Now().Nanosecond()

	timeUse := timeEnd - timeStart

	log.Println("nano time use: ", timeUse)

	log.Println("us time use: ", timeUse/1000)
	log.Println("ms time use: ", timeUse/1e6)

	ld := logger.LogDetail{
		Cost:       int64(timeUse),
		ReqID:      "",
		ReqMethod:  "test",
		ReqPath:    "",
		ClientAddr: "",
		RespCode:   0,
		RespError:  "",
	}

	ldFields := ld.AddFields()

	logger.Log.Info("", ldFields...)
}
