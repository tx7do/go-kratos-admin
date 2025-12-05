package gorm

import (
	"kratos-admin/app/admin/service/internal/data/gorm/models"

	"github.com/tx7do/go-crud/gorm"
)

var migrates []interface{}

func init() {
	RegisterMigrates()
}

func RegisterMigrates() {
	gorm.RegisterMigrateModels(
		&models.AdminLoginLog{},
		&models.AdminLoginRestriction{},
		&models.AdminOperationLog{},
		&models.ApiResource{},
		&models.Department{},
		&models.DictEntry{},
		&models.DictType{},
		&models.File{},
		&models.InternalMessage{},
		&models.InternalMessageCategory{},
		&models.InternalMessageRecipient{},
		&models.Language{},
		&models.Menu{},
		&models.Organization{},
		&models.Position{},
		&models.Role{},
		&models.RoleApi{},
		&models.RoleDept{},
		&models.RoleMenu{},
		&models.RoleOrg{},
		&models.RolePosition{},
		&models.Task{},
		&models.Tenant{},
		&models.User{},
		&models.UserCredential{},
		&models.UserPosition{},
		&models.UserRole{},
	)
}
