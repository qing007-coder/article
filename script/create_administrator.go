package script

import (
	"article/pkg/constant"
	"article/pkg/model"
	"article/pkg/rules"
	"article/pkg/tools"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateAdministrator(db *gorm.DB, enforcer *rules.Enforcer) error {
	id := tools.CreateID()
	password, err := bcrypt.GenerateFromPassword([]byte("qing040824"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	db.Create(&model.User{
		ID:       id,
		Account:  "1457800511",
		Password: string(password),
		Email:    "1457800511@qq.com",
		Status:   true,
	})
	return enforcer.AddGroup(id, constant.ADMINISTRATOR)
}
