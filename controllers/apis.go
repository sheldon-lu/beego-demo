package controllers

import (
	"beego-demo/models"
)

var (
	M              = new(models.M)
	roleModel      = new(models.RbacRole)
	apiModel       = new(models.RbacApi)
	userModel      = new(models.User)
	accessModel    = new(models.RbacAccess)
	resultTreeNode = []models.RbacApi{}
)

type ApisController struct {
	BaseController
}

func (self *ApisController) List() {
	_, action := self.GetControllerAndAction()
	self.CheckRbacAll(action)
	var result []models.RbacApi
	_, err := M.Object().QueryTable("rbac_api").Filter("level__in", []int{1, 2}).All(&result, "Id", "Name", "Title", "Level")
	if err != nil {

	}
	self.Data["ParentNode"] = result
	self.Data["LeftNavResult"] = self.BaseController.GetSession("LeftNavResult")
	self.TplName = "user/apis.html"
}

func (self *ApisController) ListJson() {
	_, action := self.GetControllerAndAction()
	self.CheckRbacAll(action)
	var result []models.RbacApi
	n, _ := M.Object().QueryTable("rbac_api").OrderBy("id").All(&result)
	resultTreeNode = []models.RbacApi{}
	self.ajaxList("success", 0, int64(n), result)
}


func (self *ApisController) Add() {
	_, action := self.GetControllerAndAction()
	self.CheckRbacAll(action)
	// self.CheckRbacAll("Role")
	id, _ := self.GetInt("id")
	name := self.GetString("name")
	title := self.GetString("title")
	pid, _ := self.GetInt("pid")
	status, _ := self.GetInt("status")
	is_show, _ := self.GetInt("is_show")
	remark := self.GetString("remark")
	sort, _ := self.GetInt("sort")
	level, _ := self.GetInt("level")
	if id != 0 {
		result := models.RbacApi{Id: id}
		if M.Object().Read(&result) == nil {
			result.Remark = remark
			result.Status = status
			result.IsShow = is_show
			result.Pid = pid
			result.Name = name
			result.Level = level
			// result.Sort = sort
			result.Title = title
			if _, err := M.Object().Update(&result); err == nil {
				self.ajaxList("update success",0, 0, nil)
			}
			self.ajaxList("update failed",102, 0, nil)
		} else {
			self.ajaxList("update failed",102, 0, nil)
		}
	}

	result := models.RbacApi{Name: name, Pid: pid, Status: status, Remark: remark, Title: title, Sort: sort, Level: level, IsShow: is_show}
	_, err := M.Object().Insert(&result)
	if err != nil {
		self.ajaxList("insert failed", 101, 0, nil)
	}
	self.ajaxList("success", 0, 0, nil)
}

func (self *ApisController) Delete() {
	_, action := self.GetControllerAndAction()
	self.CheckRbacAll(action)
	id, _ := self.GetInt("id")
	result := models.RbacApi{Id: id}
	_, err := M.Object().Delete(&result)
	if err != nil {
		self.ajaxList("delete failed", 103, 0, nil)
	}
	self.ajaxList("delete success", 0, 0, nil)

}
