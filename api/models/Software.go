package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Software struct {
	ID               uint32  `gorm:"primary_key;auto_increment" json:"id"`
	Code             string  `gorm:"size:255;not null" json:"code"`
	Name             string  `gorm:"size:255;not null" json:"Name"`
	ZipFile          string  `gorm:"size:255;" json:"ZipFile"`
	LinkSource       string  `gorm:"size:255;" json:"LinkSource"`
	LinkPreview      string  `gorm:"size:255;" json:"LinkPreview"`
	LinkTutorial     string  `gorm:"size:255;" json:"LinkTutorial"`
	License          string  `gorm:"size:255;" json:"License"`
	Description      string  `gorm:"size:255;" json:"Description"`
	PreviewImage     string  `gorm:"size:255;" json:"PreviewImage"`
	Ebook            string  `gorm:"size:255;" json:"Ebook"`
	ProductVersion   float64 `json:"ProductVersion"`
	KategoriID       uint32
	Kategori         Kategori `gorm:"foreignKey:KategoriID"`
	VideoTutorial    []VideoTutorial
	DokumenPendukung []DokumenPendukung
	ReleaseDate      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"ReleaseDate"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Software) SaveSoftware(db *gorm.DB) (*Software, error) {
	err := db.Debug().Model(&Software{}).Create(&p).Error
	if err != nil {
		return &Software{}, err
	}
	return p, nil
}

func (p *Software) GetAllSoftwares(db *gorm.DB) (*[]Software, error) {
	Softwares := []Software{}
	err := db.Debug().Model(&Software{}).Limit(100).Preload("Kategori").Preload("VideoTutorial").Preload("DokumenPendukung").Find(&Softwares).Error
	if err != nil {
		return &[]Software{}, err
	}
	return &Softwares, nil
}

func (u *Software) GetSoftwareByID(db *gorm.DB, uid uint32) (*Software, error) {
	err := db.Debug().Model(Software{}).Where("id = ?", uid).Preload("Kategori").Preload("VideoTutorial").Preload("DokumenPendukung").Take(&u).Error
	if err != nil {
		return &Software{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Software{}, errors.New("Software Not Found")
	}
	return u, err
}

func (u *Software) DeleteASoftware(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Software{}).Where("id = ?", uid).Take(&Software{}).Delete(&Software{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (u *Software) UpdateSoftware(db *gorm.DB, uid uint32) (*Software, error) {
	u.UpdatedAt = time.Now()

	db = db.Debug().Model(&Software{}).Where("id = ?", uid).Take(&Software{}).UpdateColumns(&u)
	if db.Error != nil {
		return &Software{}, db.Error
	}
	// This is the display the updated Software
	err := db.Debug().Model(&Software{}).Where("id = ?", uid).Preload("Kategori").Preload("VideoTutorial").Preload("DokumenPendukung").Take(&u).Error
	if err != nil {
		return &Software{}, err
	}
	return u, nil
}

// func (u *Software) GetSoftwareArray(db *gorm.DB, uid uint32) (*Software, error) {
// 	err := db.Debug().Model(Software{}).Where("id = ?", uid).Preload("Kategori").Preload("VideoTutorial").Preload("DokumenPendukung").Take(&u).Error
// 	if err != nil {
// 		return &Software{}, err
// 	}
// 	if gorm.IsRecordNotFoundError(err) {
// 		return &Software{}, errors.New("Software Not Found")
// 	}
// 	return u, err
// }
