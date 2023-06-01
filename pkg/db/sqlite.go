package db

import (
	v1 "gmcm/model/v1"
	"gmcm/pkg/utils"
	"gmcm/pkg/utils/auth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var dbPath = utils.GetEnv("db.path", utils.UserHome()+"/.gmcm.db")
var client *gorm.DB

func SetupDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Duration(10) * time.Second)
	sqlDB.SetMaxIdleConns(100)

	initDatabase(db)

	return db, nil
}

func Client() *gorm.DB {
	return client
}

func SetClient(c *gorm.DB) {
	client = c
}

func initDatabase(db *gorm.DB) {
	db.AutoMigrate(&v1.Hosts{}, &v1.User{}, &v1.Status{})
	var record int64
	var u = &v1.User{}
	var s = &v1.Status{}
	var now = time.Now()
	db.Table("users").Count(&record)
	if record == 0 {
		u.ObjectMeta.SetCreateAt(now)
		u.ObjectMeta.SetUpdateAt(now)
		u.UserName = "admin"
		saltpass, _ := auth.Encrypt("edoc2")
		u.Password = saltpass

		db.Table("users").Create(&u)
	}

	db.Table("statuses").Count(&record)
	if record == 0 {
		s.InitStatus = 0
		db.Table("statuses").Create(&s)
	}
}
