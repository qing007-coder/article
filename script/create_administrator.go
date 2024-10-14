package script

import (
	"article/pkg/constant"
	"article/pkg/model"
	"article/pkg/rules"
	"article/pkg/tools"
	"gorm.io/gorm"
)

func CreateAdministrator(db *gorm.DB, enforcer *rules.Enforcer) error {
	id := tools.CreateID()

	db.Create(&model.User{
		ID:       id,
		Account:  "1457800511",
		Password: "qing040824",
		Email:    "1457800511@qq.com",
		Status:   true,
	})
	return enforcer.AddGroup(id, constant.ADMINISTRATOR)
}
