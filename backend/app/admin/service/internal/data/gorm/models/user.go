package models

import (
	"time"

	"gorm.io/datatypes"

	"github.com/tx7do/go-crud/gorm/mixin"
)

// User 对应表 sys_users，用户表
type User struct {
	mixin.AutoIncrementID

	Username      *string         `gorm:"column:username;type:varchar(255);comment:用户名;uniqueIndex:idx_sys_user_username"`
	Nickname      *string         `gorm:"column:nickname;type:varchar(255);comment:昵称"`
	Realname      *string         `gorm:"column:realname;type:varchar(255);comment:真实名字"`
	Email         *string         `gorm:"column:email;type:varchar(320);comment:电子邮箱"`
	Mobile        *string         `gorm:"column:mobile;type:varchar(255);comment:手机号码"`
	Telephone     *string         `gorm:"column:telephone;type:varchar(255);comment:座机号码"`
	Avatar        *string         `gorm:"column:avatar;type:varchar(255);comment:头像"`
	Address       *string         `gorm:"column:address;type:varchar(255);comment:地址"`
	Region        *string         `gorm:"column:region;type:varchar(255);comment:国家地区"`
	Description   *string         `gorm:"column:description;type:varchar(1023);comment:个人说明"`
	Gender        *string         `gorm:"column:gender;type:enum('SECRET','MALE','FEMALE');comment:性别;index:idx_sys_user_gender"`
	Authority     *string         `gorm:"column:authority;type:enum('SYS_ADMIN','TENANT_ADMIN','CUSTOMER_USER','GUEST');comment:授权;index:idx_sys_user_authority"`
	Status        *string         `gorm:"column:status;type:enum('ON','OFF');comment:用户状态;index:idx_sys_user_status"`
	LastLoginTime *time.Time      `gorm:"column:last_login_time;type:datetime;comment:最后一次登录的时间"`
	LastLoginIP   *string         `gorm:"column:last_login_ip;type:varchar(45);comment:最后一次登录的IP"`
	OrgID         *uint32         `gorm:"column:org_id;type:int unsigned;comment:组织ID;index:idx_sys_user_org_id"`
	DepartmentID  *uint32         `gorm:"column:department_id;type:int unsigned;comment:部门ID;index:idx_sys_user_department_id"`
	PositionID    *uint32         `gorm:"column:position_id;type:int unsigned;comment:职位ID;index:idx_sys_user_position_id"`
	WorkID        *uint32         `gorm:"column:work_id;type:int unsigned;comment:员工工号"`
	RoleIDs       *datatypes.JSON `gorm:"column:role_ids;type:json;comment:角色ID列表"`

	mixin.TimeAt
	mixin.OperatorID
	mixin.Remark
	mixin.TenantID
}

// TableName 指定表名
func (User) TableName() string {
	return "sys_users"
}
