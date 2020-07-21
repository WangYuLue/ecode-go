package controllers

import (
	myJwt "ecode/middlewares/jwt"
	"ecode/models"
	"ecode/utils/md5"
	"ecode/utils/message"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var j = myJwt.NewJWT()

// Login 登录接口
func Login(c *gin.Context) {
	// name 用户名或者邮箱
	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" || password == "" {
		message.HandelError(c, message.ErrHTTPData.BindFail)
		return
	}
	_, err := models.GetUserByName(name)
	if err != nil {
		_, err := models.GetUserByEmail(name)
		if err != nil {
			message.HandelError(c, message.ErrUser.NotFound)
			return
		}
	}

	password = md5.Md5(password)
	user, err := models.Login(name, password)
	if err != nil {
		message.HandelError(c, message.ErrUser.PasswordIncorrect)
		return
	}
	GenerateToken(c, user)
	return
}

// GenerateToken 生成令牌
func GenerateToken(c *gin.Context, user models.User) {
	claims := myJwt.CustomClaims{
		ID:    strconv.Itoa(user.ID),
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),      // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600*24*7), // 过期时间 一星期
			Issuer:    "ecode",                              //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		message.HandelError(c, message.ErrToken.GenerateFail)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"data":    user,
		"token":   token,
	})
	return
}

// UpdateToken 更新token
func UpdateToken(c *gin.Context) {
	token := c.PostForm("token")
	newToken, err := j.RefreshToken(token)
	if err != nil {
		message.HandelError(c, message.ErrToken.UpdateFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": newToken,
	})
}
