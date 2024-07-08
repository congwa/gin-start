package system

import (
	"github.com/congwa/gin-start/global"
	"github.com/gofrs/uuid/v5"
)

type SysUser struct {
	global.MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`          // 用户UUID
	Username string    `json:"userName" gorm:"index;comment:用户登录名"`       // 用户登录名
	Password string    `json:"-"  gorm:"comment:用户登录密码"`                  // 用户登录密码
	NickName string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称
	Phone    string    `json:"phone"  gorm:"comment:用户手机号"`               // 用户手机号
	Email    string    `json:"email"  gorm:"comment:用户邮箱"`                // 用户邮箱
}

func (SysUser) TableName() string {
	return "sys_users"
}

// 创建表

/*
	CREATE TABLE sys_users (
    id INT AUTO_INCREMENT PRIMARY KEY,                    -- 主键，自增ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,       -- 创建时间
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 更新时间
    deleted_at TIMESTAMP NULL DEFAULT NULL,               -- 删除时间，用于软删除
    uuid CHAR(36) NOT NULL,                               -- 用户UUID，唯一标识
    user_name VARCHAR(255) NOT NULL,                      -- 用户登录名
    password VARCHAR(255) NOT NULL,                       -- 用户登录密码
    nick_name VARCHAR(255) DEFAULT '系统用户',             -- 用户昵称，默认值为“系统用户”
    phone VARCHAR(50),                                    -- 用户手机号
    email VARCHAR(255),                                   -- 用户邮箱
    INDEX idx_uuid (uuid),                                -- uuid索引
    INDEX idx_username (user_name)                        -- 用户登录名索引
);

-- 添加列的注释
ALTER TABLE sys_users
    MODIFY uuid CHAR(36) COMMENT '用户UUID',
    MODIFY user_name VARCHAR(255) COMMENT '用户登录名',
    MODIFY password VARCHAR(255) COMMENT '用户登录密码',
    MODIFY nick_name VARCHAR(255) COMMENT '用户昵称',
    MODIFY phone VARCHAR(50) COMMENT '用户手机号',
    MODIFY email VARCHAR(255) COMMENT '用户邮箱';

*/
