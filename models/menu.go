package models

import (
	"fmt"
	"web_app/mysql"

	uuid "github.com/satori/go.uuid"
)

// Menu ...
type Menu struct {
	ID       string `json:"id"`
	Pid      string `json:"pid"`
	Path     string `json:"path"`
	MenuName string `json:"menu_name"`
	Buttons  string `json:"buttons"`
}

// TableName ...
func (Menu) TableName() string {
	return "menu"
}

// TreeNode ..
type TreeNode struct {
	ID       string     `json:"id"`
	Pid      string     `json:"pid"`
	Path     string     `json:"path"`
	MenuName string     `json:"menu_name"`
	Buttons  string     `json:"buttons"`
	Children []TreeNode `json:"children"`
}

// ListToTree ...
func ListToTree(nodes []Menu, pid string) []TreeNode {
	var temp []TreeNode
	var treeArr = nodes
	for index, item := range nodes {
		if item.Pid == pid {
			var kkk TreeNode
			kkk.ID = treeArr[index].ID
			kkk.Pid = treeArr[index].Pid
			kkk.Path = treeArr[index].Path
			kkk.MenuName = treeArr[index].MenuName
			kkk.Buttons = treeArr[index].Buttons
			kkk.Children = ListToTree(treeArr, treeArr[index].ID)
			fmt.Println(kkk)
			temp = append(temp, kkk)
		}
	}
	return temp
}

// AddMenu 添加记录
func (m *Menu) AddMenu() (err error) {
	m.ID = uuid.NewV4().String()
	err = mysql.Db.Create(&m).Error
	if err != nil {
		fmt.Println("添加失败")
	}
	return
}

// // GetUsers 获取记录
// func (u *User) GetUsers(page, pagesize int) (list []User, err error) {
// 	err = mysql.Db.Where(u).Offset((page - 1) * pagesize).Limit(pagesize).Find(&list).Error
// 	if err != nil {
// 		fmt.Println("添加失败")
// 	}
// 	return
// }

// UpdateMenu ...
func (m *Menu) UpdateMenu() (err error) {
	var old Menu
	err = mysql.Db.Where("id=?", m.ID).First(&old).Error
	if err != nil {
		return err
	}
	err = mysql.Db.Model(&old).Updates(&m).Error
	if err != nil {
		fmt.Println("更新失败")
	}
	return
}

// GetMenuTree ...
func (m *Menu) GetMenuTree() (tree []TreeNode, err error) {
	var menus []Menu
	err = mysql.Db.Find(&menus).Error
	if err != nil {
		return
	}
	tree = ListToTree(menus, "")
	fmt.Println(tree)
	return
}

// // DelUser ...
// func (u *User) DelUser() (err error) {
// 	err = mysql.Db.First(u).Error
// 	if err != nil {
// 		return err
// 	}
// 	err = mysql.Db.Delete(u).Error
// 	if err != nil {
// 		fmt.Println("删除失败")
// 	}
// 	return
// }
