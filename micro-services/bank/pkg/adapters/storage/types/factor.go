package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Factor represents the GORM entity for the Factor domain.
type Factor struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SourceService  string         `gorm:"type:varchar(255);not null"`                       // Name of the service generating the factor.
	ExternalID     uuid.UUID      `gorm:"type:uuid;not null"`                               // Unique identifier in the source service.
	BookingID      uuid.UUID      `gorm:"type:uuid;not null;index"`                         // Booking ID for factor aggregation.
	Amount         uint           `gorm:"not null"`                                         // Total amount for the factor.
	CustomerID     uuid.UUID      `gorm:"not null;index"`                                   // Reference to the customer associated with the factor.
	Status         uint8          `gorm:"not null;default:0"`                               // Status of the factor (0: Unknown, 1: Pending, etc.).
	Details        string         `gorm:"type:text"`                                        // JSON or additional metadata related to the factor.
	InstantPayment bool           `gorm:"not null;default:false"`                           // Whether the payment should be made instantly.
	IsPaid         bool           `gorm:"not null;default:false"`                           // Whether the factor has been paid.
	CreatedAt      time.Time      `gorm:"autoCreateTime"`                                   // Automatically set at creation.
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`                                   // Automatically updated.
	DeletedAt      gorm.DeletedAt `gorm:"index"`                                            // Soft delete support.
	Transactions   []Transaction  `gorm:"foreignKey:FactorID;constraint:OnDelete:SET NULL"` // One-to-many relation with transactions.
}
