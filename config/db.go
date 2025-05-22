package config

import (
	"go-web1/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func initDB() {
	dsn := Appconfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to initialize database, got error: %v", err)
	}

	sqlDb, err := db.DB()

	sqlDb.SetMaxIdleConns(Appconfig.Database.MaxIdleConns) //最大连接池数量
	sqlDb.SetMaxOpenConns(Appconfig.Database.MaxOpenCons)  // 打开数据库最大数量
	sqlDb.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("Failed to configure database, got error: %v", err)
	}

	global.Db = db
}
