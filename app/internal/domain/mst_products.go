package domain

type MstProducts struct {
	ID          uint   `gorm:"primaryKey"`
	JAN         string `gorm:"not null;varcher(13);->"`
	ProductName string `gorm:"not null;varcher(50);->"`
}
