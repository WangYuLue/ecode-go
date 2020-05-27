package jwt

import (
	"ecode/config"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	errTokenExpired     error = errors.New("令牌已过期")
	errTokenNotValidYet error = errors.New("令牌尚未激活")
	errTokenInvalid     error = errors.New("令牌不合法")
)

// Auth 中间件，检查token
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		tokenData := strings.Split(authorization, " ")
		header := tokenData[0]
		if len(tokenData) != 2 || header != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "请求头 Authorization 不合法",
			})
			c.Abort()
			return
		}
		token := tokenData[1]
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "请求未携带令牌，无权限访问",
			})
			c.Abort()
			return
		}

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == errTokenExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"msg": "令牌已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			c.Abort()
			return
		}
		idInt, _ := strconv.Atoi(claims.ID)
		// 默认前 100 个用户是管理员
		if idInt > 100 && claims.ID != c.Param("userid") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "当前令牌无权限访问该资源",
			})
			c.Abort()
			return
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}

// CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(config.JWTSignKey),
	}
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken -
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errTokenNotValidYet
			} else {
				return nil, errTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errTokenInvalid
}

// RefreshToken 刷新token
// 刷新逻辑: 前端用未过期的 token 来换取最新的 token；
//          每天最多刷新一次 token，逻辑放在 axios 拦截器中实现；
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(24 * 7 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", errTokenInvalid
}
