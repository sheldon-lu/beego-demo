package controllers

import (
	"beego-demo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	loginUser      *models.User
}

func (self *BaseController) Prepare() {
	// router url */*
	controllerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.Data["siteName"] = beego.AppConfig.String("site.name")

	_, res := self.isLogin()
	if res == false && self.actionName != "login" {
		self.redirect(self.URLFor("HomeController.Login"))
	}

	if res == true {
		self.Layout = "base/layout.html"
		self.loginUser = self.GetSession("LoginUser").(*models.User)
		self.Data["id"] = self.loginUser.Id
		self.Data["employeeNumber"] = self.loginUser.EmployeeNumber
		self.Data["username"] = self.loginUser.Username
		self.Data["realName"] = self.loginUser.RealName
		self.Data["displayName"] = self.loginUser.DisplayName
		self.Data["firstName"] = self.loginUser.FirstName
		self.Data["secondName"] = self.loginUser.SecondName
		self.Data["phone"] = self.loginUser.Phone
		self.Data["email"] = self.loginUser.Email
		self.Data["company"] = self.loginUser.Company
		self.Data["department"] = self.loginUser.Department
		self.Data["position"] = self.loginUser.Position
		self.Data["onboardTime"] = self.loginUser.OnboardTime
	}
}


func (self *BaseController) isLogin() (interface{}, bool) {
	userData := self.GetSession("LoginUser")
	if userData != nil {
		return userData, true //登入状态
	}
	return nil, false //登出状态
}


func (self *BaseController) redirect(url string) {
	self.Redirect(url, 302)
	self.StopRun()
}


// ajax返回
func (self *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}


// ajax list
func (self *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}




func (self *BaseController) SessionRbac(roleid int) {
	// result := models.User{EmployeeNumber: employeeid}
	// fmt.Println(result)
	// M.Object().Read(&result)

	roleResult := models.RbacRole{Id: roleid}

	M.Object().Read(&roleResult)
	self.SetSession("RoleInfo", roleResult)

	var accessResult []orm.Params
	M.Object().Raw("select t1.id,t1.name,t1.title,t1.sort,t1.pid,t1.level,t.role_id" +
		" from rbac_access t inner join rbac_api t1" +
		" on t.api_id=t1.id " +
		" where t1.status=1 and t.role_id='" + strconv.Itoa(roleResult.Id) + "' and t1.is_show=1 order by t1.sort,t1.pid asc").Values(&accessResult)

	// leftTreeResult = []orm.Params{}

	// tmpResult := self.TreeNodeRecursion(accessResult, 0)
	// fmt.Println(leftTreeResult)
	self.SetSession("LeftNavResult", accessResult)
}

// 每次请求一个界面接口会检查用户信息
func (self *BaseController) CheckLogin() (interface{}, bool) {
	userData := self.GetSession("LoginUser")
	// fmt.Println(userData)
	if userData == nil {
		return nil, false
	}
	return userData, true
}

func (self *BaseController) CheckRbacAll(module string) {
	// var infoResult models.User

	info, err := self.CheckLogin()
	var infoResult interface{}
	infoResult = &info
	infos, _ := infoResult.(models.User)
	// fmt.Println(ok)
	if !err {
		self.Redirect("/login", 302)
		self.StopRun()
	}
	// u := beego.AppConfig.String("superadmin")
	if infos.RoleId == 1 {
		return
	}
	//err = this.CheckRbac(module)
	//if !err {
	//	self.Redirect("/404", 302)
	//	self.StopRun()
	//}
}

func (self *BaseController) CheckRbac(module string) bool {
	tmpResult := self.GetSession("LeftNavResult")
	if tmpResult == nil {
		return false
	}
	returnResult := false
	for _, v := range tmpResult.([]orm.Params) {
		// 需要添加模糊，待优化 ~= /user/* /home/* 等
		if (v["name"].(string) == module) {
			returnResult = true
			break
		}
	}
	return returnResult
}