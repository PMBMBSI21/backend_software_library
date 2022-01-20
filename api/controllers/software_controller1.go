package controllers

import (
	"fmt"
	"net/http"
	"software_library/backend/api/models"
	"software_library/backend/api/responses"
	"software_library/backend/api/utils/formaterror"
	upload "software_library/backend/api/utils/uploadfile"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (server *Server) CreateSoftware(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// }

	software := models.Software{}

	// err = json.Unmarshal(body, &software)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	now := time.Now()
	timeUpload := now.Unix()
	nameFile := upload.RandomString(3)
	code := fmt.Sprintf("%d%s", timeUpload, nameFile)

	software.Code = code
	software.Name = r.FormValue("Name")
	kategori_id, _ := strconv.ParseUint(r.FormValue("KategoriID"), 10, 32)
	software.KategoriID = uint32(kategori_id)
	software.ZipFile, _ = upload.UploadFile(w, r, "ZipFile", code)
	software.LinkSource = r.FormValue("LinkSource")
	software.LinkPreview = r.FormValue("LinkPreview")
	software.LinkTutorial = r.FormValue("LinkTutorial")
	software.License = r.FormValue("License")
	software.Description = r.FormValue("Description")
	software.PreviewImage, _ = upload.UploadFile(w, r, "PreviewImage", code)
	software.Ebook, _ = upload.UploadFile(w, r, "Ebook", code)

	productVersion, _ := strconv.ParseFloat(r.FormValue("ProductVersion"), 64)
	software.ProductVersion = productVersion

	softwareCreated, err := software.SaveSoftware(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, softwareCreated.ID))
	responses.JSON(w, http.StatusCreated, softwareCreated)
}

func (server *Server) GetSoftwares(w http.ResponseWriter, r *http.Request) {
	Software := models.Software{}

	Softwares, err := Software.GetAllSoftwares(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, Softwares)
}

func (server *Server) GetSoftware(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	Software := models.Software{}
	SoftwareGotten, err := Software.GetSoftwareByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, SoftwareGotten)
}

func (server *Server) DeleteSoftware(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	Software := models.Software{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = Software.DeleteASoftware(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) UpdateSoftware(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	software := models.Software{}
	// err = json.Unmarshal(body, &software)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	software.Name = r.FormValue("Name")
	kategori_id, _ := strconv.ParseUint(r.FormValue("KategoriID"), 10, 32)
	software.KategoriID = uint32(kategori_id)
	// software.ZipFile, _ = upload.UploadFile(w, r, "ZipFile", code)
	software.LinkSource = r.FormValue("LinkSource")
	software.LinkPreview = r.FormValue("LinkPreview")
	software.LinkTutorial = r.FormValue("LinkTutorial")
	software.License = r.FormValue("License")
	software.Description = r.FormValue("Description")
	// software.PreviewImage, _ = upload.UploadFile(w, r, "PreviewImage", code)
	// software.Ebook, _ = upload.UploadFile(w, r, "Ebook", code)

	productVersion, _ := strconv.ParseFloat(r.FormValue("ProductVersion"), 64)
	software.ProductVersion = productVersion

	updatedSoftware, err := software.UpdateSoftware(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedSoftware)
}
