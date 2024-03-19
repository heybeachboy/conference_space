package bootstrap

import (
	"ConferenceSpace/config"
	"ConferenceSpace/db/mysql"
	"ConferenceSpace/logger"
	"ConferenceSpace/protocol/ws"
)

func BootInit(path string) error {
	if err := config.InitConfig(path); err != nil {
		logger.ErrorF("Init load config error")
		return err
	}

	if err := mysql.InitMysql(); err != nil {
		logger.ErrorF("Init create mysql conn error")
		return err
	}

	go ws.HubService.Run()
	logger.InfoF("Init boot successful!!")
	return nil
}
