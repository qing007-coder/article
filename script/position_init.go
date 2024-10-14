package script

import (
	"article/pkg/constant"
	"article/pkg/rules"
)

func PositionInit(enforcer rules.Enforcer) {
	enforcer.AddPolicy(constant.ADMINISTRATOR, constant.ARTICLE, constant.JUDEG)
}
