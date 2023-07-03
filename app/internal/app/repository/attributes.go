package repository

type Attribute struct {
	ID    uint   `gorm:"primaryKey"`
	JAN   int    `gorm:"not null;->"`
	Value string `gorm:"not null;varcher(50);->"`
}
