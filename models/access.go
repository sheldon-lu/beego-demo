package models

type RbacAccess struct {
	Id     int `orm:"pk"`
	RoleId int
	ApiId  int
	Level  int
	Module string
}

type RbacApi struct {
	Id     int `orm:"pk"`
	Name   string
	Title  string
	Status int
	Remark string
	Sort   int
	Pid    int
	Level  int
	IsShow int
}

type RbacRole struct {
	Id       int           `orm:"pk"`
	Name     string
	Pid      int
	Status   int
	Remark   string
}