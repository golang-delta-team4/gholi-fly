package types

import (
	"time"

	"github.com/google/uuid"
)

type Staff struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;index"`
	AgencyID  uuid.UUID `gorm:"type:uuid;index"`
	WalletID  uuid.UUID `gorm:"type:uuid;index"`
	Stock     int
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
