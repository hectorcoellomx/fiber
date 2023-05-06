package database

import (
	"github.com/hectorcoellomx/fiber/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB(cfg config.Config) (*gorm.DB, error) {
	dsn := cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

/*

// SQL Server (:9930)

import (
	"github.com/hectorcoellomx/fiber/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func OpenDB(cfg config.Config) (*gorm.DB, error) {

	// dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", cfg.User, cfg.Password, cfg.Server, cfg.Port, cfg.Database)
	// dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"

	dsn := "sqlserver://"+ cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "?database=" + cfg.DBName
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}

*/
