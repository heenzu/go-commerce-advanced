package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID               string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ParentID         string `gorm:"size:36;index"`
	User             User
	UserID           string `gorm:"size:36;index"`
	ProductImages    []ProductImage
	Categories       []Category      `gorm:"many2many:product_categories;"`
	Name             string          `gorm:"size:255"`
	Slug             string          `gorm:"size:255"`
	Price            decimal.Decimal `gorm:"type:decimal(16,2);"`
	Stock            int
	ShortDescription string          `gorm:"type:text"`
	Description      string          `gorm:"type:text"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

func (p *Product) GetProducts(db *gorm.DB, perPage int, page int) (*[]Product, int64, error) {
	var err error
	var products []Product
	var count int64

	err = db.Debug().Model(&Product{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * perPage

	err = db.Debug().Model(&products).Order("created_at desc").Limit(perPage).Offset(int(offset)).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return &products, count, nil
}

func (p *Product) FindBySlug(db *gorm.DB, slug string) (*Product, error) {
	var err error
	var product Product

	err = db.Debug().Preload("ProductImages").Model(&Product{}).Where("slug = ?", slug).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}