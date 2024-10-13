package rules

import (
	"BackEnd/Model"
	"BackEnd/Utils"
	"gorm.io/gorm"
)

type Enforcer struct {
	db *gorm.DB
}

func NewEnforcer(db *gorm.DB) *Enforcer {

	return &Enforcer{
		db: db,
	}
}

func (e *Enforcer) AddGroup(uid, policy string) error {
	// 判断有没有这个policy
	var count int64
	e.db.Model(&Model.UserRole{}).Where("type = ? AND v1 = ?", POLICY, policy).Count(&count)
	if count == 0 {
		return NOMATCHINGPOLICY
	}

	e.db.Create(&Model.UserRole{
		ID:   Utils.CreateID(),
		Type: GROUP,
		V1:   uid,
		V2:   policy,
	})

	return nil
}

func (e *Enforcer) AddPolicy(policy, source, action string) {
	e.db.Create(&Model.UserRole{
		ID:   Utils.CreateID(),
		Type: POLICY,
		V1:   policy,
		V2:   source,
		V3:   action,
	})
}

func (e *Enforcer) Enforce(uid, source, action string) error {
	var role Model.UserRole
	if err := e.db.Where("type = ? AND v2 = ? AND v3 = ?", POLICY, source, action).First(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return NOMATCHINGPOLICY
		} else {
			return OTHERERROR
		}
	}

	var count int64
	e.db.Model(&Model.UserRole{}).Where("type = ? AND v1 = ? AND v2 = ?", GROUP, uid, role.V1).Count(&count)
	if count == 0 {
		return INSUFFICIENTPERMISSIONS
	}

	return nil
}

func (e *Enforcer) RemovePolicy(policy, source, action string) error {
	if err := e.db.Where("type = ? AND v1 = ? AND v2 = ? AND v3 = ?", POLICY, policy, source, action).Delete(&Model.UserRole{}).Error; err != nil {
		return err
	}

	return nil
}

func (e *Enforcer) RemoveGroup(uid, policy string) error {
	if err := e.db.Where("type = ? AND v1 = ? AND v2 = ?", GROUP, uid, policy).Delete(&Model.UserRole{}).Error; err != nil {
		return err
	}

	return nil
}
