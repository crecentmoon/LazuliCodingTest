package domain

type Maker struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null;varcher(255)"`
}
