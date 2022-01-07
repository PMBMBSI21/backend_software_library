package upload

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

func UploadFile(w http.ResponseWriter, r *http.Request, value_name string, folder_name string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}

	uploadedFile, handler, err := r.FormFile(value_name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}

	now := time.Now()
	timeUpload := now.Unix()
	nameFile := RandomString(10)
	filename := fmt.Sprintf("%d-%s%s", timeUpload, nameFile, filepath.Ext(handler.Filename))

	createDirectory := filepath.Join(dir, "uploads", folder_name)
	if _, err := os.Stat(createDirectory); os.IsNotExist(err) {
		fmt.Println("ga ADAA")
		os.Mkdir(createDirectory, 0777)
	}
	fileLocation := filepath.Join(dir, "uploads", folder_name, filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}

	databaseValue := fmt.Sprintf("%s%s/uploads/%s/%s", os.Getenv("HOST_NAME"), os.Getenv("FILE_PORT"), folder_name, filename)

	return databaseValue, nil
}

func RandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
