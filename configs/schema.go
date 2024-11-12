package configs

import (
	"basic-gin-app/internal/entity"
	"basic-gin-app/internal/utils"
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func GetDb() *gorm.DB {
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dbPassword := decryptSecretKey()
	fmt.Println("password: ", dbPassword)
	dsn := utils.Cfg["database"].DbUsername + ":" + dbPassword + "@tcp" + "(" + utils.Cfg["database"].DbHost + ":" + utils.Cfg["database"].DbPort + ")/" + utils.Cfg["database"].DbName + "?parseTime=true"

	sqlDB, err := sql.Open("mysql", dsn)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(time.Hour)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Error("Error connecting to database : error=%v\n", err)
		return nil
	}
	var models = []interface{}{&entity.Resource{}}
	err = db.AutoMigrate(models...)
	if err != nil {
		log.Error("Error auto-migrating :", err)
	}

	return db
}

func decryptSecretKey() string {
	decryptedPassword := utils.DecryptAes(utils.Cfg["database"].DbPassword, utils.Cfg["database"].DbEncryptionPassword)
	return decryptedPassword
}
