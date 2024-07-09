package system

import (
	"errors"
	"fmt"

	"github.com/congwa/gin-start/global"
	"github.com/congwa/gin-start/model/system"
	"github.com/congwa/gin-start/utils"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser

	err = global.DB.Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}

	return &user, err
}

func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	err = global.DB.Create(&u).Error

	return u, err
}

// resetPassword 修改用户密码
func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = global.DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}
