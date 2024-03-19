package mysql

import (
	"ConferenceSpace/config"
	"ConferenceSpace/db/model"
	"ConferenceSpace/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {

}

var dbConn *gorm.DB
var err error

func InitMysql() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.MysqlDb.Username,
		config.Config.MysqlDb.Password,
		config.Config.MysqlDb.Host,
		config.Config.MysqlDb.Port,
		"mysql")
	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s default charset utf8 COLLATE utf8_general_ci;", config.Config.MysqlDb.DbName)
	err = dbConn.Exec(sql).Error
	if err != nil {
		Close()
		return err
	}
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.MysqlDb.Username,
		config.Config.MysqlDb.Password,
		config.Config.MysqlDb.Host,
		config.Config.MysqlDb.Port,
		config.Config.MysqlDb.DbName)
	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err = dbConn.AutoMigrate(&model.User{}, model.Room{}); err != nil {
		logger.ErrorF("mysql auto migrate error : %s", err.Error())
	}
	return nil
}

func Close() error {
	db, err2 := dbConn.DB()
	if err2 != nil {
		return err2
	}
	return db.Close()
}
