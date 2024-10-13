package model

type UserRole struct {
	ID   string `gorm:"primaryKey"`
	Type string
	V1   string
	V2   string
	V3   string
}
