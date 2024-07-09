package utils

import (
	"errors"
	"time"

	"github.com/congwa/gin-start/global"
	"github.com/congwa/gin-start/model/system/request"
	jwt "github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SigningKey []byte
}

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.Config.JWT.SigningKey),
	}
}

// CreateClaims 用于根据基础声明 baseClaims 创建一个 JWT 的自定义声明 CustomClaims
// JWT 是一个 JWT 结构体指针
// baseClaims 是请求的基础声明
// 返回值是 request.CustomClaims 类型，代表 JWT 的自定义声明

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	bf, _ := ParseDuration(global.Config.JWT.BufferTime)
	ep, _ := ParseDuration(global.Config.JWT.ExpiresTime)
	claim := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf),
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"gin"},                   // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间 7天  配置文件
			Issuer:    global.Config.JWT.Issuer,                  // 签名的发行者
		},
	}
	return claim
}

// 创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token  TODO: 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	// 解析旧的 token
	token, err := jwt.ParseWithClaims(oldToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	// 检查解析是否成功
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return "", ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return "", ErrTokenNotValidYet
			} else {
				return "", ErrTokenInvalid
			}
		}
	}

	// 旧 token 有效，创建新 token
	if token != nil && token.Valid {
		return j.CreateToken(claims)
	}

	return "", ErrTokenInvalid
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid

	} else {
		return nil, ErrTokenInvalid
	}
}
