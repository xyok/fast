package db

import (
	"fmt"
	"strings"
	"time"
	"{{ .AppName }}/lib/log"
	"{{ .AppName }}/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)
var db *gorm.DB

func createDB(url string) error {
	dbURL := url
	slashIndex := strings.LastIndex(dbURL, "/")
	dsn := dbURL[:slashIndex+1]
	dbName := dbURL[slashIndex+1:]
	if strings.Contains(dbName, "?") {
		dbName = strings.Split(dbName, "?")[0]
	}
	log.Info("dsn[%v] db[%s]", dsn, dbName)
	dsn = fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", dsn)
	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	createSQL := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS   `%s` CHARACTER SET utf8mb4;",
		dbName,
	)

	log.Info("%s", createSQL)

	err = db.Exec(createSQL).Error
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

//Setup db
func Setup() error {
	URL := conf.Database.URL
	if err := createDB(URL); err != nil {
		return err
	}

	_db, err := gorm.Open(mysql.Open(URL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})

	if err != nil {
		log.Error("%v", err)
		log.Fatal("db url %v", URL)
	}

	sqlDB, err := _db.DB()
	if err != nil {
		log.Error("%v", err)
		panic(err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(conf.Database.MaxIdleConn)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(conf.Database.MaxIdleConn)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Second * 1800)

	if conf.Database.Debug {
		log.Info("db show SQL --> true")
		_db.Config.Logger = _db.Config.Logger.LogMode(logger.Info)
	}

	db = _db

	log.Debug("db setup")
	return nil
}

//GetDB db
func GetDB() *gorm.DB {
	return db
}

//CloseDB close
func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Error("%v", err)
	}
	sqlDB.Close()
}
