package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"software_library/backend/api/models"
	"software_library/backend/api/responses"
	"software_library/backend/api/utils/formaterror"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) CreateKategori(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	Kategori := models.Kategori{}

	// Kategori.Name = r.FormValue("Name")

	err = json.Unmarshal(body, &Kategori)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	KategoriCreated, err := Kategori.SaveKategori(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, KategoriCreated.ID))
	responses.JSON(w, http.StatusCreated, KategoriCreated)
}

func (server *Server) GetKategoris(w http.ResponseWriter, r *http.Request) {
	Kategori := models.Kategori{}

	Kategoris, err := Kategori.GetAllKategoris(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, Kategoris)
}

func (server *Server) GetKategori(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	Kategori := models.Kategori{}
	KategoriGotten, err := Kategori.GetKategoriByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, KategoriGotten)
}

func (server *Server) DeleteKategori(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	Kategori := models.Kategori{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = Kategori.DeleteAKategori(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) UpdateKategori(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	Kategori := models.Kategori{}
	err = json.Unmarshal(body, &Kategori)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedKategori, err := Kategori.UpdateKategori(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedKategori)
}
