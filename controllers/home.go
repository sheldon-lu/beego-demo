package controllers

import (
	"beego-demo/common"
	"beego-demo/common/hjwt"
	"beego-demo/models"
	"fmt"
	"strconv"

	"gopkg.in/ldap.v3"
)

type HomeController struct {
	BaseController
}

func (self *HomeController) Logout() {
	self.DelSession("LoginUser")
	self.redirect(self.URLFor("HomeController.Login"))
}

func (self *HomeController) Login() {
	if self.Ctx.Request.Method == "GET" {
		self.TplName = "home/login.html"
	} else {
		username := self.GetString("username")
		password := self.GetString("password")
		// result := models.User{}
		token := hjwt.GenToken()
		fmt.Println(token)
		// entry, res := self.LoginCheck(username, password)
		LoginUser, res := self.LoginCheck1(username, password)
		resultroleid, _ := self.Permission(username)
		if res == true {
			// LoginUser := self.UpdateUser(entry, password)
			self.SetSession("LoginUser", LoginUser)
			self.Ctx.SetCookie("Authorization", token, 360, "/*")
			self.SessionRbac(resultroleid)
			self.redirect(self.URLFor("HomeController.Index"))
		} else {
			self.redirect(self.URLFor("HomeController.Login"))
		}
		//self.SetSession("LoginUser", "true")
		//self.SessionRbac(resultroleid)
		//self.Ctx.SetCookie("Authorization", token, 360, "/")
		//self.redirect(self.URLFor("HomeController.Index"))
	}
}

// LDAP
func (self *HomeController) LoginCheck(username, password string) (*ldap.Entry, bool) {
	userInfo, err := common.LdapAuth(username, password)
	fmt.Println(userInfo, err)
	if userInfo != nil {
		//验证成功
		return userInfo, true
	} else {
		//验证失败
		return nil, false
	}
}

// COMMON
func (self *HomeController) LoginCheck1(username, password string) (*models.User, bool) {
	userInfo := models.UserGetByUsername(username)
	fmt.Println(password)
	fmt.Println(userInfo.Password)
	if userInfo != nil && userInfo.Password == password {
		//验证成功
		return userInfo, true
	} else {
		//验证失败
		return nil, false
	}

}

func (self *HomeController) UpdateUser(entry *ldap.Entry, password string) *models.User {
	//判断数据库中是否有username，有则修改，没有则添加
	user := models.UserGetByUsername(entry.GetAttributeValue("uid"))
	user.EmployeeNumber, _ = strconv.Atoi(entry.GetAttributeValue("employeeNumber"))
	user.Password = entry.GetAttributeValue("userPassword")
	user.RealName = entry.GetAttributeValue("cn")
	user.DisplayName = entry.GetAttributeValue("displayName")
	user.FirstName = entry.GetAttributeValue("givenName")
	user.SecondName = entry.GetAttributeValue("sn")
	user.Phone = entry.GetAttributeValue("mobile")
	user.Email = entry.GetAttributeValue("mail")
	user.Company = entry.GetAttributeValue("o")
	user.Position = entry.GetAttributeValue("title")
	user.OnboardTime = entry.GetAttributeValue("description")
	if user.Username != "" {
		// 更新操作
		if err := user.Update(); err != nil {
			panic(err)
		}
		return user
	} else {
		//添加操作
		user.Username = entry.GetAttributeValue("uid")
		// fmt.Println(user)
		if _, err := models.UserAdd(user); err != nil {
			panic(err)
		}
		return user
	}
}

func (self *HomeController) Index() {
	self.Data["LeftNavResult"] = self.GetSession("LeftNavResult")
	self.TplName = "home/index.html"
}

func (self *HomeController) Permission(username string) (roleid int, bo bool) {
	userIf := models.UserGetByUsername(username)
	roleid = userIf.RoleId
	// fmt.Println(userIf)
	// fmt.Println(userIf.RoleId)
	if roleid == 1 {
		return roleid, true
	} else {
		return roleid, false
	}
}
