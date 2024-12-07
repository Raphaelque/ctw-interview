package model

import (
	"os"
	"strings"
	"time"

	"ctw-interview/common"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// var SQL_DSN=root:123456@tcp(mysql:3306)/new-api
func InitDB() (err error) {
	db, err := chooseDB("SQL_DSN")
	if err == nil {

		if common.DebugEnabled {
			db = db.Debug()
		}
		DB = db
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		// 数据库连接池设置
		sqlDB.SetMaxIdleConns(common.GetEnvOrDefault("SQL_MAX_IDLE_CONNS", 100))
		sqlDB.SetMaxOpenConns(common.GetEnvOrDefault("SQL_MAX_OPEN_CONNS", 1000))
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(common.GetEnvOrDefault("SQL_MAX_LIFETIME", 60)))

		// 自动创建数据库
		common.SysLog("database migration started")
		err = migrateDB()
		return err
	} else {
		common.FatalLog(err)
	}
	return err
}

func chooseDB(env string) (*gorm.DB, error) {
	dsn := os.Getenv(env)
	if dsn != "" {
		// Use MySQL
		common.SysLog("using MySQL as database")
		// check parseTime
		if !strings.Contains(dsn, "parseTime") {
			if strings.Contains(dsn, "?") {
				dsn += "&parseTime=true"
			} else {
				dsn += "?parseTime=true"
			}
		}
		common.UsingMySQL = true
		return gorm.Open(mysql.Open(dsn), &gorm.Config{
			PrepareStmt: true, // precompile SQL
		})
	}
	if !strings.Contains(dsn, "parseTime") {
		if strings.Contains(dsn, "?") {
			dsn += "&parseTime=true"
		} else {
			dsn += "?parseTime=true"
		}
	}
	// Use SQLite 如果没有设置SQL_DSN环境变量，则使用SQLite作为数据库
	common.SysLog("SQL_DSN not set, using SQLite as database")
	common.UsingSQLite = true
	return gorm.Open(sqlite.Open(common.SQLitePath), &gorm.Config{
		PrepareStmt: true, // precompile SQL
	})
}

func migrateDB() error {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&Task{})
	if err != nil {
		return err
	}
	common.SysLog("database migrated")
	return err
}
