package types

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrUnableToDeleteSuperAdmin = errors.New("can not delete SuperAdmin role")
)

type Role struct {
	gorm.Model
	UUID        uuid.UUID
	Name        string       `gorm:"unique"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
	UserRole []UserRole
}

type UserRole struct {
	ID        uint `gorm:"primarykey"`
	UserID    uint
	User      *User `gorm:"foreignKey:UserID"`
	RoleID    uint
	Role      *Role `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
}

func (r *Role) BeforeDelete(tx *gorm.DB) (err error) {
	if r.Name == "SuperAdmin" {
		return ErrUnableToDeleteSuperAdmin
	}
	return
}

func (r *Role) AfterDelete(tx *gorm.DB) (error) {
    err := tx.Clauses(clause.Returning{}).Where("user_roles.role_id = ?", r.ID).Delete(&UserRole{}).Error
	if err != nil {
		return err
	}
    err = tx.Table("role_permissions").Where("role_id = ?", r.ID).Delete(nil).Error
    return err
}
