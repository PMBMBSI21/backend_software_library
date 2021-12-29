package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type DokumenPendukung struct {
	ID          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"size:255;not null" json:"Title"`
	File        string `gorm:"size:255;not null" json:"Url"`
	Description string `gorm:"size:255;not null" json:"Description"`
	SoftwareID  uint32
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *DokumenPendukung) SaveDokumenPendukung(db *gorm.DB) (*DokumenPendukung, error) {
	err := db.Debug().Model(&DokumenPendukung{}).Create(&p).Error
	if err != nil {
		return &DokumenPendukung{}, err
	}
	return p, nil
}

func (p *DokumenPendukung) GetAllDokumenPendukungs(db *gorm.DB) (*[]DokumenPendukung, error) {
	DokumenPendukungs := []DokumenPendukung{}
	err := db.Debug().Model(&DokumenPendukung{}).Limit(100).Find(&DokumenPendukungs).Error
	if err != nil {
		return &[]DokumenPendukung{}, err
	}
	return &DokumenPendukungs, nil
}

func (u *DokumenPendukung) GetDokumenPendukungByID(db *gorm.DB, uid uint32) (*DokumenPendukung, error) {
	err := db.Debug().Model(DokumenPendukung{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &DokumenPendukung{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &DokumenPendukung{}, errors.New("video Tutorial Not Found")
	}
	return u, err
}

func (u *DokumenPendukung) DeleteADokumenPendukung(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&DokumenPendukung{}).Where("id = ?", uid).Take(&DokumenPendukung{}).Delete(&DokumenPendukung{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (u *DokumenPendukung) UpdateDokumenPendukung(db *gorm.DB, uid uint32) (*DokumenPendukung, error) {
	u.UpdatedAt = time.Now()

	db = db.Debug().Model(&DokumenPendukung{}).Where("id = ?", uid).Take(&DokumenPendukung{}).UpdateColumns(&u)
	if db.Error != nil {
		return &DokumenPendukung{}, db.Error
	}
	// This is the display the updated DokumenPendukung
	err := db.Debug().Model(&DokumenPendukung{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &DokumenPendukung{}, err
	}
	return u, nil
}
