package script

import (
	"article/pkg/rules"
	"article/pkg/tools"
)

func CreateAdministrator(enforcer *rules.Enforcer) {
	id := tools.CreateID()

	enforcer.AddGroup()
}
