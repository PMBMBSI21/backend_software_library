package seed

import (
	"fmt"
	"log"
	"os"
	"software_library/backend/api/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func getFile(foldername string, filename string) string {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}

	fileLocation := fmt.Sprintf("%s%s/uploads/%s/%s", os.Getenv("HOST_NAME"), os.Getenv("HOST_PORT"), foldername, filename)

	return fileLocation
}

var users = []models.User{
	models.User{
		Name:     "Administrator",
		Email:    "admin@gmail.com",
		Password: "admin",
		Level:    2,
	},
	models.User{
		Name:     "User 1",
		Email:    "user@gmail.com",
		Password: "user",
		Level:    1,
	},
}

var softwares = []models.Software{
	models.Software{
		Code:           "EXAMPLE01",
		Name:           "Pembelajaran Python",
		ZipFile:        getFile("EXAMPLE01", "example-software.zip"),
		LinkSource:     "https://github.com/",
		LinkPreview:    "https://github.com/",
		LinkTutorial:   "https://www.youtube.com/",
		License:        "Open Source",
		Description:    "Example data software",
		PreviewImage:   getFile("EXAMPLE01", "example-image.png"),
		Ebook:          getFile("EXAMPLE01", "example-pdf.pdf"),
		ProductVersion: 3.0,
	},
	models.Software{
		Code:           "EXAMPLE02",
		Name:           "Projek Perkebunan",
		ZipFile:        getFile("EXAMPLE02", "example-software.zip"),
		LinkSource:     "https://github.com/",
		LinkPreview:    "https://github.com/",
		LinkTutorial:   "https://www.youtube.com/",
		License:        "Open Source",
		Description:    "Example data software",
		PreviewImage:   getFile("EXAMPLE02", "example-image.png"),
		Ebook:          getFile("EXAMPLE02", "example-pdf.pdf"),
		ProductVersion: 5.0,
	},
}

var video = []models.VideoTutorial{
	models.VideoTutorial{
		Title:       "Example Video 1",
		Url:         "https://www.youtube.com/",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy",
	},
	models.VideoTutorial{
		Title:       "Example Video 2",
		Url:         "https://www.youtube.com/",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy",
	},
}

var dokumen = []models.DokumenPendukung{
	models.DokumenPendukung{
		Name:        "Example Document 1",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy",
	},
	models.DokumenPendukung{
		Name:        "Example Document 2",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}, &models.Software{}, &models.VideoTutorial{}, &models.DokumenPendukung{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Software{}, &models.VideoTutorial{}, &models.DokumenPendukung{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i, _ := range softwares {
		err = db.Debug().Model(&models.Software{}).Create(&softwares[i]).Error
		if err != nil {
			log.Fatalf("cannot seed softwares table: %v", err)
		}

		video[i].SoftwareID = softwares[i].ID

		err = db.Debug().Model(&models.VideoTutorial{}).Create(&video[i]).Error
		if err != nil {
			log.Fatalf("cannot seed video table: %v", err)
		}

		dokumen[i].SoftwareID = softwares[i].ID
		dokumen[i].File = softwares[i].Ebook

		err = db.Debug().Model(&models.DokumenPendukung{}).Create(&dokumen[i]).Error
		if err != nil {
			log.Fatalf("cannot seed dokumen table: %v", err)
		}
	}
}
