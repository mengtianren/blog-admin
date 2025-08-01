package config

import (
	"blog-admin/global"
	"blog-admin/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbInit() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		Config.Database.User,
		Config.Database.Password,
		Config.Database.Host,
		Config.Database.Port,
		Config.Database.DBName,
		Config.Database.Charset,
		Config.Database.ParseTime,
		Config.Database.Loc,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	global.DB = db

	// sqlDb, err := db.DB()

	// if err != nil {
	// 	return nil, err
	// }
	// sqlDb.SetMaxOpenConns()
	if Config.App.Env == "dev" {
		err1 := db.AutoMigrate(&models.User{}, &models.Blog{}, &models.Role{}, &models.Comment{})
		if err1 != nil {
			fmt.Printf("数据库迁移失败:%v", err1)
		}
	}

}
