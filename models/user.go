package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id             int
	EmployeeNumber int    //employeeNumber
	Username       string //uid
	RealName       string //cn
	DisplayName    string //displayName
	FirstName      string //givenName
	SecondName     string //sn
	Password       string //userPassword
	Phone          string //mobile
	Email          string //mail
	Company        string //o
	Department     string //departmentNumber
	Position       string //title
	OnboardTime    string //onboardTime
	RoleId         int
}

func (a *User) TableName() string {
	return TableName("user")
}

// 获取列表
func UserGetList(page, pageSize int, filters ...interface{}) ([]*User, int64) {
	offset := (page - 1) * pageSize
	list := make([]*User, 0)
	query := orm.NewOrm().QueryTable(TableName("user"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

// cy_user 根据username查询
func UserGetByUsername(username string) *User {
	res := new(User)
	_ = orm.NewOrm().QueryTable(TableName("user")).Filter("username", username).One(res)
	return res
}


func UserGetById(id int) (*User, error) {
	r := new(User)
	err := orm.NewOrm().QueryTable(TableName("user")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// cy_user_info 更新数据
func (a *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

// cy_user_info 添加数据
func UserAdd(a *User) (int64, error) {
	return orm.NewOrm().Insert(a)
}

//cy_user_info 删除数据
func UserDelById(id int) bool {
	o := orm.NewOrm()
	_ = o.Using("default")
	//根据ID得到用户模型
	if num, err := o.Delete(&User{Id: id}); err == nil {
		fmt.Println("删除影响的行数:")
		fmt.Println(num)
		return true
	}else{
		return false
	}
}
