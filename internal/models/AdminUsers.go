/*
 * @Description:用户相关model
 * @Author: gphper
 * @Date: 2021-07-04 11:58:45
 */
package models

import (
	"github/gphper/ginadmin/pkg/comment"
	"time"
)

type AdminUsers struct {
	BaseModle
	Uid       uint   `gorm:"primary_key;auto_increment"`
	GroupName string `gorm:"size:20;comment:'用户组名称'"`
	Username  string `gorm:"size:100;comment:'用户名'"`
	Nickname  string `gorm:"size:100;comment:'姓名'"`
	Password  string `gorm:"size:200;comment:'密码'"`
	Phone     string `gorm:"size:20;comment:'手机号'"`
	LastLogin string `gorm:"size:30;comment:'最后登录ip地址'"`
	Salt      string `gorm:"size:32;comment:'密码盐'"`
	ApiToken  string `gorm:"size:32;comment:'用户登录凭证'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AdminUserIndexReq struct {
	Nickname  string `form:"nickname"`
	CreatedAt string `form:"created_at"`
}

type AdminUserSaveReq struct {
	Username  string   `form:"username" label:"用户名" binding:"required"`
	Password  string   `form:"password"`
	Nickname  string   `form:"nickname" label:"姓名" binding:"required"`
	Phone     string   `form:"phone"`
	GroupName []string `form:"groupname[]" label:"用户组" binding:"required"`
	Uid       uint     `form:"uid"`
}

type AdminUserEditPassReq struct {
	Uid         int    `form:"id"`
	OldPassword string `form:"old_password" label:"原始密码" binding:"required"`
	NewPassword string `form:"new_password" label:"新密码" binding:"required"`
	SubPassword string `form:"sub_password" label:"确认密码" binding:"required"`
}

func (au *AdminUsers) TableName() string {
	return "admin_users"
}

func (au *AdminUsers) FillData() {
	//初始化管理员
	salt := comment.RandString(6)
	passwordSalt := comment.Encryption("111111", salt)
	adminUser := AdminUsers{
		Uid:       1,
		GroupName: "superadmin",
		Username:  "admin",
		Nickname:  "管理员",
		Password:  passwordSalt,
		Phone:     "",
		LastLogin: "",
		Salt:      salt,
		ApiToken:  "",
	}
	Db.Save(&adminUser)
}
