package models

import (
	"fmt"
	"time"
	"web_app/mysql"

	uuid "github.com/satori/go.uuid"
)

// User 用户模型
// @Description 用户模型定义
type User struct {
	ID        string    `json:"id"`                           //id
	Username  string    `json:"username" binding:"required"`  //用户名
	Password  string    `json:"password" binding:"required"`  //密码
	RoleName  string    `json:"role_name" binding:"required"` //角色名称
	RoleID    string    `json:"role_id" binding:"required"`   //角色id
	CreatedAt time.Time `json:"created_at" `                  //创建时间，不用传
	UpdateAt  time.Time `json:"updated_at" `                  //更新时间，不用传
}

// TableName 数据库
func (User) TableName() string {
	return "user"
}

// AddUser 添加记录
func (u *User) AddUser() (err error) {
	u.ID = uuid.NewV4().String()
	u.UpdateAt = time.Now()
	u.CreatedAt = time.Now()
	err = mysql.Db.Create(&u).Error
	if err != nil {
		fmt.Println("添加失败")
	}
	return
}

// GetUsers 获取记录
func (u *User) GetUsers(page, pagesize int) (list []User, err error) {
	err = mysql.Db.Where(u).Offset((page - 1) * pagesize).Limit(pagesize).Find(&list).Error
	if err != nil {
		fmt.Println("添加失败")
	}
	return
}

// Login ...
func (u *User) Login() (err error) {
	err = mysql.Db.Where(u).Find(u).Error
	if err != nil {
		fmt.Println("用户和密码不正确")
	}
	return
}

// UpdateUser ...
func (u *User) UpdateUser() (err error) {
	var oldUser User
	u.UpdateAt = time.Now()
	err = mysql.Db.Where("id=?", u.ID).First(&oldUser).Error
	if err != nil {
		return err
	}
	u.CreatedAt = oldUser.CreatedAt
	err = mysql.Db.Model(&oldUser).Updates(&u).Error
	if err != nil {
		fmt.Println("更新失败")
	}
	return
}

// DelUser ...
func (u *User) DelUser() (err error) {
	err = mysql.Db.First(u).Error
	if err != nil {
		return err
	}
	err = mysql.Db.Delete(u).Error
	if err != nil {
		fmt.Println("删除失败")
	}
	return
}
