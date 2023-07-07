package repository

import "encoding/json"

type MakerModel struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null;varchar(255)"`
}

type BrandModel struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null;varchar(255)"`
}

type ProductModel struct {
	JAN             int64                 `gorm:"primaryKey"`
	ProductName     string                `gorm:"not null;varchar(255)"`
	Maker           MakerModel            `gorm:"foreignKey:MakerID"`
	Brand           BrandModel            `gorm:"foreignKey:BrandID"`
	Attribute       AttributeModel        `gorm:"foreignKey:JAN"`
	DescriptionTags []DescriptionTagModel `gorm:"foreignKey:JAN"`
	ReviewTags      []ReviewTagModel      `gorm:"foreignKey:JAN"`
}

type AttributeModel struct {
	ID    uint            `gorm:"primaryKey"`
	JAN   int64           `gorm:"not null"`
	Value json.RawMessage `gorm:"not null"`
}

type DescriptionTagModel struct {
	ID  uint   `gorm:"primaryKey"`
	JAN int64  `gorm:"not null"`
	Tag string `gorm:"not null;varchar(255)"`
}

type ReviewTagModel struct {
	ID  uint   `gorm:"primaryKey"`
	JAN int64  `gorm:"not null"`
	Tag string `gorm:"not null;varchar(255)"`
}
