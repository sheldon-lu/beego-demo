package controllers

import (
	"beego-demo/models"
)

type RoleController struct {
	BaseController
}

func (self *RoleController) List() {
	_, action := self.GetControllerAndAction()
	self.CheckRbacAll(action)
	var result []models.RbacApi
	M.Object().QueryTable(apiModel).OrderBy("sort").All(&result)
	resultTreeNode = []models.RbacApi{}
	// result = self.TreeNodeRecursion(result, 0)
	self.Data["LeftNavResult"] = self.BaseController.GetSession("LeftNavResult")
	self.Data["NodeResult"] = result
	self.TplName = "user/role.html"
}

func (self *RoleController) ListJson() {
	_, action := self.GetControllerAndAction()
	self.CheckRbacAll(action)
	var result []models.RbacRole
	n, _ := M.Object().QueryTable("rbac_role").All(&result)
	self.ajaxList("success", 0,  int64(n), result)
}


func (self *RoleController) Add() {
	_, action := self.GetControllerAndAction()
	self.CheckRbacAll(action)
	id, _ := self.GetInt("id")
	name := self.GetString("name")
	pid, _ := self.GetInt("pid")
	status, _ := self.GetInt("status")
	remark := self.GetString("remark")
	if id != 0 {
		result := models.RbacRole{Id: id}
		if M.Object().Read(&result) == nil {
			result.Remark = remark
			result.Status = status
			result.Pid = pid
			result.Name = name
			if _, err := M.Object().Update(&result); err == nil {
				self.ajaxList("update success", 0, 0, nil)
			}
			self.ajaxList("update failed", 102, 0, nil)
		} else {
			self.ajaxList("update failed", 102, 0, nil)
		}
	}

	result := models.RbacRole{Name: name, Pid: pid, Status: status, Remark: remark}
	_, err := M.Object().Insert(&result)
	if err != nil {
		self.ajaxList("insert failed", 101, 0, nil)
	}
	self.ajaxList("success", 0, 0, nil)
}

func (self *RoleController) Delete() {
	_, action := self.GetControllerAndAction()
	self.CheckRbacAll(action)
	id, _ := self.GetInt("id")
	result := models.RbacRole{Id: id}
	_, err := M.Object().Delete(&result)
	if err != nil {
		self.ajaxList("delete failed", 103, 0, nil)
	}
	self.ajaxList("delete success", 0, 0, nil)

}