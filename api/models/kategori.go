package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Kategori struct {
	ID         uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name       string `gorm:"size:255;not null" json:"Title"`
	SoftwareID []Software
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Kategori) SaveKategori(db *gorm.DB) (*Kategori, error) {
	err := db.Debug().Model(&Kategori{}).Create(&p).Error
	if err != nil {
		return &Kategori{}, err
	}
	return p, nil
}

func (p *Kategori) GetAllKategoris(db *gorm.DB) (*[]Kategori, error) {
	Kategoris := []Kategori{}
	err := db.Debug().Model(&Kategori{}).Limit(100).Preload("SoftwareID").Find(&Kategoris).Error
	if err != nil {
		return &[]Kategori{}, err
	}
	return &Kategoris, nil
}

func (u *Kategori) GetKategoriByID(db *gorm.DB, uid uint32) (*Kategori, error) {
	err := db.Debug().Model(Kategori{}).Where("id = ?", uid).Preload("SoftwareID").Take(&u).Error
	if err != nil {
		return &Kategori{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Kategori{}, errors.New("video Tutorial Not Found")
	}
	return u, err
}

func (u *Kategori) DeleteAKategori(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Kategori{}).Where("id = ?", uid).Take(&Kategori{}).Delete(&Kategori{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (u *Kategori) UpdateKategori(db *gorm.DB, uid uint32) (*Kategori, error) {
	u.UpdatedAt = time.Now()

	db = db.Debug().Model(&Kategori{}).Where("id = ?", uid).Take(&Kategori{}).UpdateColumns(&u)
	if db.Error != nil {
		return &Kategori{}, db.Error
	}
	// This is the display the updated Kategori
	err := db.Debug().Model(&Kategori{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Kategori{}, err
	}
	return u, nil
}
