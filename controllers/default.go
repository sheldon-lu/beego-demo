package controllers

type MainController struct {
	BaseController
}

func (self *MainController) Get() {
	// 每次渲染左侧导航需要这个获取session
	//self.Data["LeftNavResult"] = self.BaseController.GetSession("LeftNavResult")
	self.TplName = "home/index.html"
}
