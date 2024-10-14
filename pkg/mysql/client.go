package mysql

import (
	"article/pkg/config"
	"article/pkg/errors"
	"article/pkg/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewClient(conf *config.GlobalConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		conf.Mysql.Name,
		conf.Mysql.Password,
		conf.Mysql.Address,
		conf.Mysql.Port,
		conf.Mysql.Database,
		conf.Mysql.Conf,
	)

	// 初始化数据库时的高级配置
	client, err := gorm.Open(mysql.New(mysql.Config{
		// 数据库句柄
		DSN: dsn,
		// 数据库默认字符串长度
		DefaultStringSize: 256,
	}), &gorm.Config{})

	if err != nil {
		return nil, errors.DBInitFailed
	}

	if err := client.AutoMigrate(
		&model.UserRole{},
		&model.ArticleJudgeRecord{},
		&model.User{},
	); err != nil {
		return nil, errors.AutoMigrateFailed
	}

	return client, nil
}
