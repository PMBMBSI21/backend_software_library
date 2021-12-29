package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type VideoTutorial struct {
	ID          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `gorm:"size:255;not null" json:"Title"`
	Url         string `gorm:"size:255;not null" json:"Url"`
	Description string `gorm:"size:255;not null" json:"Description"`
	SoftwareID  uint32
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *VideoTutorial) SaveVideoTutorial(db *gorm.DB) (*VideoTutorial, error) {
	err := db.Debug().Model(&VideoTutorial{}).Create(&p).Error
	if err != nil {
		return &VideoTutorial{}, err
	}
	return p, nil
}

func (p *VideoTutorial) GetAllVideoTutorials(db *gorm.DB) (*[]VideoTutorial, error) {
	VideoTutorials := []VideoTutorial{}
	err := db.Debug().Model(&VideoTutorial{}).Limit(100).Find(&VideoTutorials).Error
	if err != nil {
		return &[]VideoTutorial{}, err
	}
	return &VideoTutorials, nil
}

func (u *VideoTutorial) GetVideoTutorialByID(db *gorm.DB, uid uint32) (*VideoTutorial, error) {
	err := db.Debug().Model(VideoTutorial{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &VideoTutorial{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &VideoTutorial{}, errors.New("video Tutorial Not Found")
	}
	return u, err
}

func (u *VideoTutorial) DeleteAVideoTutorial(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&VideoTutorial{}).Where("id = ?", uid).Take(&VideoTutorial{}).Delete(&VideoTutorial{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (u *VideoTutorial) UpdateVideoTutorial(db *gorm.DB, uid uint32) (*VideoTutorial, error) {
	u.UpdatedAt = time.Now()

	db = db.Debug().Model(&VideoTutorial{}).Where("id = ?", uid).Take(&VideoTutorial{}).UpdateColumns(&u)
	if db.Error != nil {
		return &VideoTutorial{}, db.Error
	}
	// This is the display the updated VideoTutorial
	err := db.Debug().Model(&VideoTutorial{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &VideoTutorial{}, err
	}
	return u, nil
}
