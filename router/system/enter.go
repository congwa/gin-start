package system

type RouterGroup struct {
	BaseRouter
	UserRouter
}

var RouterGroupApp = new(RouterGroup)
