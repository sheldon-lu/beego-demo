package common

import (
	"beego-demo/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/ldap.v3"
	"strconv"
)

var (
	ldapHost   = beego.AppConfig.String("ldap.host")
	ldapPort   = beego.AppConfig.DefaultInt("ldap.port", 389)
	ldapBindDN = beego.AppConfig.String("ldap.binddn")
	ldapBindPW = beego.AppConfig.String("ldap.bindpw")
	ldapBaseDN = beego.AppConfig.String("ldap.basedn")
	ldapFilter = beego.AppConfig.String("ldap.filter")
)

func LdapAuth(username, password string) (*ldap.Entry, error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapHost, ldapPort))
	if err != nil {
		return nil, err
	}
	defer l.Close()

	err = l.Bind(ldapBindDN, ldapBindPW)
	if err != nil {
		return nil, err
	}

	searchRequest := ldap.NewSearchRequest(
		ldapBaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0,
		false, fmt.Sprintf(ldapFilter, username),
		nil, nil,
	)

	sr, _ := l.Search(searchRequest)
	if len(sr.Entries) == 0 {
		return nil, nil
	} else {
		err = l.Bind(sr.Entries[0].DN, password)
		if err != nil {
			return nil, err
		}
		fmt.Println(sr.Entries[0])
		return sr.Entries[0], nil
	}
}

// 修改用户信息
func LdapModifyInfo(employee_number int, username, real_name, display_name, first_name, second_name, phone, email, company, department, position, onboard_time string) {
	l, _ := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapHost, ldapPort))
	defer l.Close()
	_ = l.Bind(ldapBindDN, ldapBindPW)

	modify := ldap.NewModifyRequest(fmt.Sprintf("uid=%s,%s", username, ldapBaseDN), nil)
	modify.Replace("employeeNumber", []string{strconv.Itoa(employee_number)})
	modify.Replace("cn", []string{real_name})
	modify.Replace("displayName", []string{display_name})
	modify.Replace("givenName", []string{first_name})
	modify.Replace("sn", []string{second_name})
	modify.Replace("mobile", []string{phone})
	modify.Replace("mail", []string{email})
	modify.Replace("o", []string{company})
	modify.Replace("departmentNumber", []string{department})
	modify.Replace("title", []string{position})
	modify.Replace("description", []string{onboard_time})

	_ = l.Modify(modify)
}

// 修改用户密码
func LdapModifyPwByUser(username, oldPassword, newPassword, newPasswordConfirm string) (bool, error) {
	l, _ := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapHost, ldapPort))
	defer l.Close()
	err := l.Bind(fmt.Sprintf("uid=%s,%s", username, ldapBaseDN), oldPassword)
	if err != nil {
		return false, err
	}

	if newPassword != newPasswordConfirm {
		err := errors.New("两次密码不同")
		return false, err
	}

	pwModifyRequest := ldap.NewPasswordModifyRequest("", oldPassword, newPassword)
	pwModifyResponse, err := l.PasswordModify(pwModifyRequest)
	if err != nil {
		return false, err
	}
	_ = pwModifyResponse.GeneratedPassword
	return true, nil
}

// 删除用户
func LdapDelByUsername(username string) {
	l, _ := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapHost, ldapPort))
	defer l.Close()
	_ = l.Bind(ldapBindDN, ldapBindPW)

	del := ldap.NewDelRequest(fmt.Sprintf("uid=%s,%s", username, ldapBaseDN), nil)

	_ = l.Del(del)
}

// 添加用户
func LdapAdd(user models.User) (bool, error) {
	l, _ := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapHost, ldapPort))
	defer l.Close()
	_ = l.Bind(ldapBindDN, ldapBindPW)

	add := ldap.NewAddRequest(fmt.Sprintf("uid=%s,%s", user.Username, ldapBaseDN), nil)
	add.Attribute("employeeNumber", []string{strconv.Itoa(user.EmployeeNumber)})
	add.Attribute("userPassword", []string{user.Username})
	add.Attribute("cn", []string{user.RealName})
	add.Attribute("displayName", []string{user.DisplayName})
	add.Attribute("givenName", []string{user.FirstName})
	add.Attribute("sn", []string{user.SecondName})
	add.Attribute("mobile", []string{user.Phone})
	add.Attribute("mail", []string{user.Email})
	add.Attribute("o", []string{user.Company})
	add.Attribute("departmentNumber", []string{user.Department})
	add.Attribute("title", []string{user.Position})
	add.Attribute("description", []string{user.OnboardTime})
	add.Attribute("objectClass", []string{"inetOrgPerson"})

	fmt.Print(add)
	fmt.Print(string(user.EmployeeNumber))

	err := l.Add(add)
	if err != nil {
		return false, err
	} else {
		passwordModifyRequest := ldap.NewPasswordModifyRequest(fmt.Sprintf("uid=%s,%s", user.Username, ldapBaseDN), "", user.Username)
		_, err = l.PasswordModify(passwordModifyRequest)
		return true, err
	}
}

// 修改密码（管理员）
func LdapModifyPwByAdmin(username, password string) (bool, error) {
	l, _ := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapHost, ldapPort))
	defer l.Close()
	_ = l.Bind(ldapBindDN, ldapBindPW)

	passwordModifyRequest := ldap.NewPasswordModifyRequest(fmt.Sprintf("uid=%s,%s", username, ldapBaseDN), "", password)
	_, err := l.PasswordModify(passwordModifyRequest)

	if err != nil {
		return false, err
	}
	return true, nil
}