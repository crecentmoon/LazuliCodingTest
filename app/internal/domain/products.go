package domain

type Product struct {
	JAN             int              `gorm:"primaryKey"`
	ProductName     string           `gorm:"not null;varcher(255)"`
	Makers          []Maker          `gorm:"foreignKey:MakerID"`
	Brands          []Brand          `gorm:"foreignKey:BrandID"`
	Attributes      []Attribute      `gorm:"foreignKey:JAN"`
	DescriptionTags []DescriptionTag `gorm:"foreignKey:JAN"`
	ReviewTags      []ReviewTag      `gorm:"foreignKey:JAN"`
}
