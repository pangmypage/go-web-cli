package models

import (
	"fmt"
	"time"
	"web_app/mysql"

	uuid "github.com/satori/go.uuid"
)

// Role 角色
type Role struct {
	ID        string    `json:"id"`
	RoleName  string    `json:"role_name" binding:"required"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"created_at" `
	UpdateAt  time.Time `json:"updated_at" `
	AuthList  []Auth    `json:"auth_list" gorm:"ForeignKey:role_id;AssociationForeignKey:id" ` //设置外键关联会自动创建从表的数据
}

// TableName ...
func (r *Role) TableName() string {
	return "role"
}

// Auth ...
type Auth struct {
	ID         string `json:"id"`
	RoleID     string `json:"role_id"`
	Path       string `json:"path" binding:"required"`
	MenuName   string `json:"menu_name" binding:"required"`
	ButtonRole string `json:"button_role"`
}

// TableName ...
func (r *Auth) TableName() string {
	return "auth"
}

// AddRole ...
func (r *Role) AddRole() (err error) {
	r.CreatedAt = time.Now()
	r.UpdateAt = time.Now()
	r.ID = uuid.NewV4().String()
	for i := range r.AuthList {
		r.AuthList[i].RoleID = r.ID
		r.AuthList[i].ID = uuid.NewV4().String()
	}
	fmt.Println(r)
	err = mysql.Db.Create(r).Error
	if err != nil {
		fmt.Println("新增失败", err)
		return
	}

	return
}

// UpdateRole ...
func (r *Role) UpdateRole() (err error) {
	r.UpdateAt = time.Now()
	err = mysql.Db.Where("role_id=?", r.ID).Delete(&Auth{}).Error
	if err != nil {
		fmt.Println(err, "删除旧数据失败")
		return
	}
	for i := range r.AuthList {
		r.AuthList[i].RoleID = r.ID
		r.AuthList[i].ID = uuid.NewV4().String()
	}
	err = mysql.Db.Save(r).Error
	if err != nil {
		fmt.Println("更新失败", err)
		return
	}
	return

}

// DelRole ...
func (r *Role) DelRole(id string) (err error) {
	err = mysql.Db.Where("id=?", id).Delete(&Role{}).Error
	if err != nil {
		return err
	}
	err = mysql.Db.Where("role_id=?", id).Delete(&Auth{}).Error
	if err != nil {
		fmt.Println("删除失败")
		return err
	}
	return
}

// GetRoles 获取角色信息
func (r *Role) GetRoles(page, pagesize int) (list []Role, err error) {
	err = mysql.Db.Where(r).Offset((page - 1) * pagesize).Limit(pagesize).Preload("AuthList").Find(&list).Error
	if err != nil {
		fmt.Println("添加失败")
	}
	return
}
