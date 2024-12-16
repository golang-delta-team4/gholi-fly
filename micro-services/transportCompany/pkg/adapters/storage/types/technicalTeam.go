package types

import "gorm.io/gorm"

type TechnicalTeam struct {
	gorm.Model
	Name string
}
