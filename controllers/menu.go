package controllers

import (
	"net/http"
	"web_app/models"
	"web_app/utils"

	"github.com/gin-gonic/gin"
)

// AddMenu ...
func AddMenu(c *gin.Context) {
	var menu models.Menu
	err := c.ShouldBind(&menu)
	if err != nil {
		c.JSON(http.StatusOK, utils.ResFaild("参数错误"))
		return
	}
	if menu.ID == "" {
		err = menu.AddMenu()
		if err != nil {
			c.JSON(http.StatusOK, utils.ResFaild("添加用户失败"))
			return
		}
	} else {
		err = menu.UpdateMenu()
		if err != nil {
			c.JSON(http.StatusOK, utils.ResFaild("更新用户失败"))
			return
		}
	}
	c.JSON(http.StatusOK, utils.ResSuccess("ok", "操作成功"))
}

// GetMenuTree 获取菜单tree
func GetMenuTree(c *gin.Context) {
	menu := models.Menu{}
	tree, err := menu.GetMenuTree()

	if err != nil {
		c.JSON(http.StatusOK, utils.ResFaild("查询失败"))
		return
	}
	c.JSON(http.StatusOK, utils.ResSuccess(tree, "操作成功"))
}
