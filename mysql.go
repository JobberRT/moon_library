package moon_library

import (
	"github.com/jobberrt/moon_library/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitMySQL(url string) error {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}
	if err := db.AutoMigrate(
		&model.Node{},
		&model.User{},
		&model.Document{},
		&model.Plan{},
		&model.Setting{},
		&model.Ticket{},
		&model.UserGroup{},
		&model.DNSTemplate{},
		&model.FallbackTemplate{},
		&model.InboundTemplate{},
		&model.OutboundTemplate{},
		&model.RoutingTemplate{},
		&model.StreamSettingTemplate{},
	); err != nil {
		return err
	}
	DB = db
	return nil
}
