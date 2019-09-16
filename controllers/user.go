package controllers

import (
	"beego-demo/common"
	"beego-demo/models"
	"fmt"
	"strconv"
	"strings"
)

type UserController struct {
	BaseController
}

func (self *UserController) Edit() {
	self.Data["LeftNavResult"] = self.BaseController.GetSession("LeftNavResult")
	self.TplName = "user/edit.html"
}

func (self *UserController) Modify() {
	resetPwd := self.GetString("reset_pwd")
	if resetPwd == "1" {
		// 修改密码
		if res, err := self.ModifyPwByUser(); res {
			self.ModifyInfo()
			self.ajaxMsg("", MSG_OK)
		} else {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	} else {
		self.ModifyInfo()
		self.ajaxMsg("", MSG_OK)
	}
}

func (self *UserController) ModifyPwByUser() (bool, error) {
	userId, _ := self.GetInt("id")
	user, _ := models.UserGetById(userId)

	oldPassword := strings.TrimSpace(self.GetString("password_old"))
	newPassword := strings.TrimSpace(self.GetString("password_new1"))
	newPasswordConfirm := strings.TrimSpace(self.GetString("password_new2"))
	user.Password = newPassword
	res, err := common.LdapModifyPwByUser(user.Username, oldPassword, newPassword, newPasswordConfirm)
	if !res {
		return false, err
	} else {
		user.Password = newPassword
		_ = user.Update()
		return true, nil
	}
}

func (self *UserController) ModifyInfo() {
	userId, _ := self.GetInt("id")
	user, _ := models.UserGetById(userId)
	user.Id = userId
	user.EmployeeNumber, _ = strconv.Atoi(strings.TrimSpace(self.GetString("employee_number")))
	user.RealName = strings.TrimSpace(self.GetString("real_name"))
	user.DisplayName = strings.TrimSpace(self.GetString("display_name"))
	user.SecondName = strings.TrimSpace(self.GetString("second_name"))
	user.FirstName = strings.TrimSpace(self.GetString("first_name"))
	user.Phone = strings.TrimSpace(self.GetString("phone"))
	user.Email = strings.TrimSpace(self.GetString("email"))
	user.Company = strings.TrimSpace(self.GetString("company"))
	user.Department = strings.TrimSpace(self.GetString("department"))
	user.Position = strings.TrimSpace(self.GetString("position"))
	user.OnboardTime = strings.TrimSpace(self.GetString("onboard_time"))
	fmt.Println(user)
	common.LdapModifyInfo(user.EmployeeNumber, user.Username, user.RealName, user.DisplayName, user.FirstName,
		user.SecondName, user.Phone, user.Email, user.Company, user.Department, user.Position, user.OnboardTime)

	_ = user.Update()
}

func (self *UserController) List() {
	self.Data["LeftNavResult"] = self.BaseController.GetSession("LeftNavResult")
	self.TplName = "user/list.html"
}

func (self *UserController) All() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	realName := strings.TrimSpace(self.GetString("realName"))

	//查询条件
	filters := make([]interface{}, 0)
	if realName != "" {
		filters = append(filters, "real_name__icontains", realName)
	}
	result, count := models.UserGetList(page, limit, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["employee_number"] = v.EmployeeNumber
		row["username"] = v.Username
		row["real_name"] = v.RealName
		row["display_name"] = v.DisplayName
		row["first_name"] = v.FirstName
		row["second_name"] = v.SecondName
		row["phone"] = v.Phone
		row["email"] = v.Email
		row["company"] = v.Company
		row["department"] = v.Department
		row["position"] = v.Position
		row["onboard_time"] = v.OnboardTime
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *UserController) EditByAdmin() {
	self.ModifyInfo()
	self.TplName = "user/list.html"
	self.ajaxMsg("修改成功", MSG_OK)
}

func (self *UserController) ModifyPwByAdmin() {
	userId, _ := self.GetInt("id")
	user, _ := models.UserGetById(userId)

	username := strings.TrimSpace(self.GetString("username"))
	password1 := strings.TrimSpace(self.GetString("password1"))
	password2 := strings.TrimSpace(self.GetString("password2"))
	fmt.Println(username, password1, password2)

	if password1 != password2 {
		self.ajaxMsg("两次输入不同", MSG_ERR)
	} else {
		if _, err := common.LdapModifyPwByAdmin(username, password1); err != nil {
			self.ajaxMsg("修改失败", MSG_ERR)
		} else {
			_ = user.Update()
			self.ajaxMsg("修改成功", MSG_OK)
		}
	}
}

func (self *UserController) Del() {
	userId, _ := self.GetInt("id")
	userName := self.GetString("username")

	common.LdapDelByUsername(userName)
	models.UserDelById(userId)
	self.TplName = "user/list.html"

	self.ajaxMsg("删除成功", MSG_OK)
}

func (self *UserController) Add() {
	user := new(models.User)

	user.EmployeeNumber, _ = strconv.Atoi(strings.TrimSpace(self.GetString("employee_number")))
	user.Username = strings.TrimSpace(self.GetString("username"))
	user.RealName = strings.TrimSpace(self.GetString("real_name"))
	user.DisplayName = strings.TrimSpace(self.GetString("display_name"))
	user.SecondName = strings.TrimSpace(self.GetString("second_name"))
	user.FirstName = strings.TrimSpace(self.GetString("first_name"))
	user.Phone = strings.TrimSpace(self.GetString("phone"))
	user.Email = strings.TrimSpace(self.GetString("email"))
	user.Company = strings.TrimSpace(self.GetString("company"))
	user.Department = strings.TrimSpace(self.GetString("department"))
	user.Position = strings.TrimSpace(self.GetString("position"))
	user.OnboardTime = strings.TrimSpace(self.GetString("onboard_time"))
	fmt.Println(user)
	fmt.Println("company=", user.Company)
	_, err := common.LdapAdd(*user)
	if err != nil {
		panic(err)
		self.TplName = "user/list.html"
		self.ajaxMsg("添加失败", MSG_ERR)
	} else {
		if _, err := models.UserAdd(user); err != nil {
			panic(err)
		}
		self.TplName = "user/list.html"
		self.ajaxMsg("添加成功", MSG_OK)
	}
}
