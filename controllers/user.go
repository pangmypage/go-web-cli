package controllers

import (
	"fmt"
	"net/http"
	"web_app/models"
	"web_app/utils"

	"github.com/gin-gonic/gin"
)

type queryUser struct {
	Username string `json:"username"`  //用户名
	RoleName string `json:"role_name"` //角色名称
	PageSize int    `json:"pagesize"`  //每页数量
	Page     int    `json:"page"`      //页码
}
type loginQuery struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type delUser struct {
	ID string `json:"id" binding:"required"`
}

// AddUser 添加用户接口
// @Summary 添加用户接口
// @Description 添加用户
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "用户令牌"
// @Param params body models.User false "查询参数"
// @Router /user/save [post]
func AddUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, utils.ResFaild("参数错误"))
		return
	}
	if user.ID == "" {
		err = user.AddUser()
		if err != nil {
			c.JSON(http.StatusOK, utils.ResFaild("添加用户失败"))
			return
		}
	} else {
		err = user.UpdateUser()
		if err != nil {
			c.JSON(http.StatusOK, utils.ResFaild("更新用户失败"))
			return
		}
	}
	c.JSON(http.StatusOK, utils.ResSuccess("ok", "操作成功"))
}

// DelUsers 删除用户接口
// @Summary 删除用户接口
// @Description 删除用户
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "用户令牌"
// @Param params body delUser false "查询参数"
// @Router /user/del [post]
func DelUsers(c *gin.Context) {
	var query delUser
	err := c.ShouldBind(&query)
	if err != nil {
		c.JSON(http.StatusOK, utils.ResFaild("参数错误"))
		return
	}
	user := models.User{
		ID: query.ID,
	}
	err = user.DelUser()
	if err != nil {
		c.JSON(http.StatusOK, utils.ResFaild("删除失败"))
		return
	}
	c.JSON(http.StatusOK, utils.ResSuccess("ok", "操作成功"))
}

// GetUsers 获取用户接口
// @Summary 获取用户接口
// @Description 删除用户
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "用户令牌"
// @Param params body queryUser false "参数"
// @Success 200 {string} json ""
// @Router /user/getlist [post]
func GetUsers(c *gin.Context) {
	var query queryUser
	err := c.ShouldBind(&query)
	if err != nil {
		c.JSON(http.StatusOK, utils.ResFaild("参数错误"))
		return
	}
	user := models.User{
		RoleName: query.RoleName,
		Username: query.Username,
	}

	list, err := user.GetUsers(query.Page, query.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, utils.ResFaild("查询失败"))
		return
	}
	c.JSON(http.StatusOK, utils.ResSuccess(list, "操作成功"))
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var query loginQuery
	err := c.ShouldBind(&query)
	if err != nil {
		c.JSON(http.StatusOK, utils.ResFaild("请输入用户名和密码"))
		return
	}
	user := models.User{
		Password: query.Password,
		Username: query.Username,
	}
	err = user.Login()
	if err != nil {
		c.JSON(http.StatusOK, utils.ResFaild("用户名或密码不正确"))
		return
	}
	token, err := utils.CreatToken(user.ID)
	if err != nil {
		fmt.Println("token生成失败")
		return
	}
	c.JSON(http.StatusOK, utils.ResSuccess(token, "登录成功"))
}
