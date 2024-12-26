package types

import (
	"time"

	"github.com/google/uuid"
)

type Agency struct {
	ID               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name             string
	OwnerID          uuid.UUID `gorm:"type:uuid;index"` // References the Staff ID
	Owner            Staff     `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	WalletID         uuid.UUID `gorm:"type:uuid;index"`
	ProfitPercentage float64
	Factors          []Factor `gorm:"foreignKey:AgencyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // One-to-many with Factor
	Staffs           []Staff  `gorm:"foreignKey:AgencyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // One-to-many with Staff
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `gorm:"index"`
}
