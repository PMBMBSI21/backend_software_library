package seed

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"software_library/backend/api/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func getFile(foldername string, filename string) string {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}

	fileLocation := fmt.Sprintf("http://%s%s/uploads/%s/%s", os.Getenv("HOST_NAME"), os.Getenv("FILE_PORT"), foldername, filename)

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
		KategoriID:     uint32(rand.Intn(5-1) + 1),
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
		KategoriID:     uint32(rand.Intn(5-1) + 1),
		ProductVersion: 5.0,
	},
	models.Software{
		Code:           "EXAMPLE03",
		Name:           "Sistem Akuntansi Perbankan",
		ZipFile:        getFile("EXAMPLE03", "example-software.zip"),
		LinkSource:     "https://github.com/",
		LinkPreview:    "https://github.com/",
		LinkTutorial:   "https://www.youtube.com/",
		License:        "Open Source",
		Description:    "Example data software",
		PreviewImage:   getFile("EXAMPLE03", "example-image.png"),
		Ebook:          getFile("EXAMPLE03", "example-pdf.pdf"),
		KategoriID:     uint32(rand.Intn(5-1) + 1),
		ProductVersion: 5.0,
	},
	models.Software{
		Code:           "EXAMPLE04",
		Name:           "Build a Bulk File Rename Tool With Python and PyQt",
		ZipFile:        getFile("EXAMPLE04", "example-software.zip"),
		LinkSource:     "https://github.com/",
		LinkPreview:    "https://github.com/",
		LinkTutorial:   "https://www.youtube.com/",
		License:        "Open Source",
		Description:    "Example data software",
		PreviewImage:   getFile("EXAMPLE04", "example-image.png"),
		Ebook:          getFile("EXAMPLE04", "example-pdf.pdf"),
		KategoriID:     uint32(rand.Intn(5-1) + 1),
		ProductVersion: 5.0,
	},
	models.Software{
		Code:           "EXAMPLE05",
		Name:           "Build a Personal Diary With Django and Python",
		ZipFile:        getFile("EXAMPLE05", "example-software.zip"),
		LinkSource:     "https://github.com/",
		LinkPreview:    "https://github.com/",
		LinkTutorial:   "https://www.youtube.com/",
		License:        "Open Source",
		Description:    "Example data software",
		PreviewImage:   getFile("EXAMPLE05", "example-image.png"),
		Ebook:          getFile("EXAMPLE05", "example-pdf.pdf"),
		KategoriID:     uint32(rand.Intn(5-1) + 1),
		ProductVersion: 5.0,
	},
	models.Software{
		Code:           "EXAMPLE06",
		Name:           "Using Pygame to Build an Asteroids Game in Python",
		ZipFile:        getFile("EXAMPLE06", "example-software.zip"),
		LinkSource:     "https://github.com/",
		LinkPreview:    "https://github.com/",
		LinkTutorial:   "https://www.youtube.com/",
		License:        "Open Source",
		Description:    "Example data software",
		PreviewImage:   getFile("EXAMPLE06", "example-image.png"),
		Ebook:          getFile("EXAMPLE06", "example-pdf.pdf"),
		KategoriID:     uint32(rand.Intn(5-1) + 1),
		ProductVersion: 5.0,
	},
	models.Software{
		Code:           "EXAMPLE07",
		Name:           "Build a Content Aggregator in Python",
		ZipFile:        getFile("EXAMPLE07", "example-software.zip"),
		LinkSource:     "https://github.com/",
		LinkPreview:    "https://github.com/",
		LinkTutorial:   "https://www.youtube.com/",
		License:        "Open Source",
		Description:    "Example data software",
		PreviewImage:   getFile("EXAMPLE07", "example-image.png"),
		Ebook:          getFile("EXAMPLE07", "example-pdf.pdf"),
		KategoriID:     uint32(rand.Intn(5-1) + 1),
		ProductVersion: 5.0,
	},
	models.Software{
		Code:           "EXAMPLE08",
		Name:           "Build a Command-Line To-Do App With Python and Typer",
		ZipFile:        getFile("EXAMPLE08", "example-software.zip"),
		LinkSource:     "https://github.com/",
		LinkPreview:    "https://github.com/",
		LinkTutorial:   "https://www.youtube.com/",
		License:        "Open Source",
		Description:    "Example data software",
		PreviewImage:   getFile("EXAMPLE08", "example-image.png"),
		Ebook:          getFile("EXAMPLE08", "example-pdf.pdf"),
		KategoriID:     uint32(rand.Intn(5-1) + 1),
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
	models.VideoTutorial{
		Title:       "Example Video 3",
		Url:         "https://www.youtube.com/",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy",
	},
}

var dokumen = []models.DokumenPendukung{
	models.DokumenPendukung{
		Title:       "Example Document 1",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy",
	},
	models.DokumenPendukung{
		Title:       "Example Document 2",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy",
	},
	models.DokumenPendukung{
		Title:       "Example Document 3",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy",
	},
	models.DokumenPendukung{
		Title:       "Example Document 4",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy",
	},
}

var kategori = []models.Kategori{
	models.Kategori{
		Name: "Website",
	},
	models.Kategori{
		Name: "Android",
	},
	models.Kategori{
		Name: "Arduino",
	},
	models.Kategori{
		Name: "Desktop",
	},
	models.Kategori{
		Name: "Ios",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}, &models.Software{}, &models.VideoTutorial{}, &models.DokumenPendukung{}, &models.Kategori{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Software{}, &models.VideoTutorial{}, &models.DokumenPendukung{}, &models.Kategori{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i, _ := range kategori {
		err = db.Debug().Model(&models.User{}).Create(&kategori[i]).Error
		if err != nil {
			log.Fatalf("cannot seed kategori table: %v", err)
		}
	}

	for i, _ := range softwares {
		err = db.Debug().Model(&models.Software{}).Create(&softwares[i]).Error
		if err != nil {
			log.Fatalf("cannot seed softwares table: %v", err)
		}

		for j, _ := range video {
			video[j].SoftwareID = softwares[i].ID
			err = db.Debug().Model(&models.VideoTutorial{}).Create(&video[j]).Error
			if err != nil {
				log.Fatalf("cannot seed video table: %v", err)
			}
			video[j].ID = 0
		}

		for k, _ := range dokumen {
			dokumen[k].SoftwareID = softwares[i].ID
			dokumen[k].FileDocument = softwares[i].Ebook

			err = db.Debug().Model(&models.DokumenPendukung{}).Create(&dokumen[k]).Error
			if err != nil {
				log.Fatalf("cannot seed dokumen table: %v", err)
			}
			dokumen[k].ID = 0
		}

	}
}
