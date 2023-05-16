package database

// SQL Server

import (
	"github.com/hectorcoellomx/fiber/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func OpenDB(cfg config.Config) (*gorm.DB, error) {

  //dsn := "sqlserver://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "?database=" + cfg.DBName
	dsn := "sqlserver://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/SQL2008?database=" + cfg.DBName + "&encrypt=disable&connection+timeout=30"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}


/*

// MySQL

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
*/
