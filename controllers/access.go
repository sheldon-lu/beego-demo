package controllers

import (
	"beego-demo/models"
	"strconv"
	"strings"
)

type AccessController struct {
	BaseController
}

func (self *AccessController) ListJson() {
	_, action := self.GetControllerAndAction()
	self.CheckRbacAll(action)
	roleid, _ := self.GetInt("role_id")
	result := []models.RbacAccess{}
	n, err := M.Object().QueryTable(accessModel).Filter("role_id", roleid).All(&result, "Id", "RoleId", "NodeId", "Level", "Module")
	if err != nil {

	}
	self.ajaxList("success", 0, int64(n), result)
}

func (self *AccessController) Add() {
	_, action := self.GetControllerAndAction()
	self.CheckRbacAll(action)
	roleid, _ := self.GetInt("role_id")
	if roleid == 0 {
		self.ajaxList("failed", 0, 0, nil)
	}
	str := self.GetString("params")
	data := []string{}
	if len(str) > 0 {
		data = strings.Split(str, ",")
	} else {
		self.ajaxList("failed",0, 0, nil)
	}
	//清空某个角色的权限，再重新添加新的
	M.Object().QueryTable(accessModel).Filter("role_id", roleid).Delete()
	result := models.RbacAccess{}
	for _, v := range data {
		tmpStr := strings.Split(v, "_")
		nodeid, _ := strconv.Atoi(tmpStr[0])
		level, _ := strconv.Atoi(tmpStr[1])
		tmpModule := ""
		if level == 1 {
			tmpModule = "项目"
		} else if level == 2 {
			tmpModule = "模块"
		} else {
			tmpModule = "操作"
		}
		result.RoleId = roleid
		result.Level = level
		result.Module = tmpModule
		result.ApiId = nodeid
		M.Object().Insert(&result)
	}
	self.ajaxList("success",0, 0, nil)
}