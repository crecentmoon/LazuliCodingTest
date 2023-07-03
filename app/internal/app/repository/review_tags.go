package repository

type ReviewTag struct {
	ID  uint   `gorm:"primaryKey"`
	JAN int    `gorm:"not null"`
	Tag string `gorm:"varcher(255)"`
}
