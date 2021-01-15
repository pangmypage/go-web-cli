package controllers

import (
	"net/http"
	"web_app/models"
	"web_app/utils"

	"github.com/gin-gonic/gin"
)

type queryRole struct {
	RoleName string `json:"role_name"`
	PageSize int    `json:"pagesize"`
	Page     int    `json:"page"`
}

type delRole struct {
	ID string `json:"id" binding:"required"`
}

// AddRole ...
func AddRole(c *gin.Context) {
	var role models.Role
	err := c.ShouldBind(&role)
	if err != nil {
		c.JSON(http.StatusOK, utils.ResFaild("参数错误"))
		return
	}
	if role.ID == "" {
		err = role.AddRole()
		if err != nil {
			c.JSON(http.StatusOK, utils.ResFaild("添加角色失败"))
			return
		}
	} else {
		err = role.UpdateRole()
		if err != nil {
			c.JSON(http.StatusOK, utils.ResFaild("更新角色失败"))
			return
		}
	}
	c.JSON(http.StatusOK, utils.ResSuccess("ok", "操作成功"))

}

// DelRoles 删除角色
func DelRoles(c *gin.Context) {
	var query delRole
	err := c.ShouldBind(&query)
	if err != nil {
		c.JSON(http.StatusOK, utils.ResFaild("参数错误"))
		return
	}
	var role models.Role
	err = role.DelRole(query.ID)
	if err != nil {
		c.JSON(http.StatusOK, utils.ResFaild("删除失败"))
		return
	}
	c.JSON(http.StatusOK, utils.ResSuccess("ok", "删除成功"))
}

// GetRoles 获取用户
func GetRoles(c *gin.Context) {
	var query queryRole
	err := c.ShouldBind(&query)
	if err != nil {
		c.JSON(http.StatusOK, utils.ResFaild("参数错误"))
		return
	}
	role := models.Role{
		RoleName: query.RoleName,
	}

	list, err := role.GetRoles(query.Page, query.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, utils.ResFaild("查询失败"))
		return
	}
	c.JSON(http.StatusOK, utils.ResSuccess(list, "操作成功"))
}
