package model

import (
	"fmt"
	"go-boss/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var Db = InitDB()

type Model struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func InitDB() *gorm.DB {
	c := config.NewConfig()
	viper := c.Viper
	username := viper.Get("mysql.username") // 账号
	password := viper.Get("mysql.password") // 密码
	host := viper.Get("mysql.host")         // 地址
	port := viper.Get("mysql.port")         // 端口
	database := viper.Get("mysql.database") // 数据库名称

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect mysql.")
	}
	migrateErr := db.AutoMigrate(&Admin{})
	if migrateErr != nil {
		panic("failed to connect migrate.")
	}
	return db
}
