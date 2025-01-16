package routers

import (
	"awesomeProject/src/routers/manager"
	"awesomeProject/src/routers/user"
)

type RouterGroup struct {
	User  user.UserRouterGroup
	Admin manager.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)
