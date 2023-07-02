package domain

type DescriptionTag struct {
	ID  uint   `gorm:"primaryKey"`
	JAN int    `gorm:"not null"`
	Tag string `gorm:"not null;varcher(255)"`
}
