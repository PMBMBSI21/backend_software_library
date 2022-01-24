package controllers

import (
	"fmt"
	"net/http"
	"software_library/backend/api/models"
	"software_library/backend/api/responses"
	"software_library/backend/api/utils/formaterror"
	upload "software_library/backend/api/utils/uploadfile"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) CreateDokumenPendukung(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// }

	DokumenPendukung := models.DokumenPendukung{}

	idForm := r.FormValue("id_software")
	uid, _ := strconv.ParseUint(idForm, 10, 32)
	Software := models.Software{}
	SoftwareById, err := Software.GetSoftwareByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	DokumenPendukung.Title = r.FormValue("Title")
	DokumenPendukung.FileDocument, _ = upload.UploadFile(w, r, "FileDocument", SoftwareById.Code)
	DokumenPendukung.Description = r.FormValue("Description")
	DokumenPendukung.SoftwareID = uint32(uid)

	// err = json.Unmarshal(body, &DokumenPendukung)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	DokumenPendukungCreated, err := DokumenPendukung.SaveDokumenPendukung(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, DokumenPendukungCreated.ID))
	responses.JSON(w, http.StatusCreated, DokumenPendukungCreated)
}

func (server *Server) GetDokumenPendukungs(w http.ResponseWriter, r *http.Request) {
	DokumenPendukung := models.DokumenPendukung{}

	DokumenPendukungs, err := DokumenPendukung.GetAllDokumenPendukungs(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, DokumenPendukungs)
}

func (server *Server) GetDokumenPendukung(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	DokumenPendukung := models.DokumenPendukung{}
	DokumenPendukungGotten, err := DokumenPendukung.GetDokumenPendukungByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, DokumenPendukungGotten)
}

func (server *Server) DeleteDokumenPendukung(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	DokumenPendukung := models.DokumenPendukung{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = DokumenPendukung.DeleteADokumenPendukung(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) UpdateDokumenPendukung(w http.ResponseWriter, r *http.Request) {

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
	DokumenPendukung := models.DokumenPendukung{}
	// err = json.Unmarshal(body, &DokumenPendukung)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	idForm := r.FormValue("id_software")
	softwareid, _ := strconv.ParseUint(idForm, 10, 32)
	Software := models.Software{}
	SoftwareById, err := Software.GetSoftwareByID(server.DB, uint32(softwareid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	DokumenPendukung.Title = r.FormValue("Title")
	DokumenPendukung.FileDocument, _ = upload.UploadFile(w, r, "FileDocument", SoftwareById.Code)
	DokumenPendukung.Description = r.FormValue("Description")
	DokumenPendukung.SoftwareID = uint32(softwareid)

	updatedDokumenPendukung, err := DokumenPendukung.UpdateDokumenPendukung(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedDokumenPendukung)
}
