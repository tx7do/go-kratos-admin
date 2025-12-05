package models

import (
	"time"

	"github.com/tx7do/go-crud/gorm/mixin"
)

type AdminLoginLog struct {
	mixin.AutoIncrementID
	mixin.CreatedAt

	LoginIP        *string    `gorm:"column:login_ip;type:varchar(64);comment:登录IP地址"`
	LoginMAC       *string    `gorm:"column:login_mac;type:varchar(128);comment:登录MAC地址"`
	LoginTime      *time.Time `gorm:"column:login_time;type:datetime;comment:登录时间"`
	UserAgent      *string    `gorm:"column:user_agent;type:text;comment:浏览器的用户代理信息"`
	BrowserName    *string    `gorm:"column:browser_name;type:varchar(128);comment:浏览器名称"`
	BrowserVersion *string    `gorm:"column:browser_version;type:varchar(128);comment:浏览器版本"`
	ClientID       *string    `gorm:"column:client_id;type:varchar(128);comment:客户端ID"`
	ClientName     *string    `gorm:"column:client_name;type:varchar(128);comment:客户端名称"`
	OSName         *string    `gorm:"column:os_name;type:varchar(128);comment:操作系统名称"`
	OSVersion      *string    `gorm:"column:os_version;type:varchar(128);comment:操作系统版本"`
	UserID         *uint32    `gorm:"column:user_id;comment:操作者用户ID"`
	Username       *string    `gorm:"column:username;type:varchar(128);comment:操作者账号名"`
	StatusCode     *int32     `gorm:"column:status_code;comment:状态码"`
	Success        *bool      `gorm:"column:success;comment:操作成功"`
	Reason         *string    `gorm:"column:reason;type:varchar(255);comment:登录失败原因"`
	Location       *string    `gorm:"column:location;type:varchar(255);comment:登录地理位置"`
}

func (AdminLoginLog) TableName() string {
	return "sys_admin_login_logs"
}
