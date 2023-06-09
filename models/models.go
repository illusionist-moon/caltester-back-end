package models

import (
	"ChildrenMath/pkg/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func init() {
	var err error

	dsn := settings.DSN

	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second * 10, // 慢SQL阈值
			LogLevel:                  logger.Info,      // 日志级别
			IgnoreRecordNotFoundError: false,            // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,             // 彩色打印
		})

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   false, // 跳过默认事务
		FullSaveAssociations:                     false,
		Logger:                                   newLogger,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		IgnoreRelationshipsWhenMigrating:         false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		TranslateError:                           false,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	})

	if err != nil {
		panic(err)
	}
}
