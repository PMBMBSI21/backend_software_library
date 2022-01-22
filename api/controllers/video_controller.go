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

func (server *Server) CreateVideoTutorial(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	VideoTutorial := models.VideoTutorial{}

	// idForm := r.FormValue("id_software")
	// uid, _ := strconv.ParseUint(idForm, 10, 32)

	// VideoTutorial.Title = r.FormValue("Title")
	// VideoTutorial.Url = r.FormValue("Url")
	// VideoTutorial.Description = r.FormValue("Description")
	// VideoTutorial.SoftwareID = uint32(uid)

	err = json.Unmarshal(body, &VideoTutorial)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	VideoTutorialCreated, err := VideoTutorial.SaveVideoTutorial(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, VideoTutorialCreated.ID))
	responses.JSON(w, http.StatusCreated, VideoTutorialCreated)
}

func (server *Server) GetVideoTutorials(w http.ResponseWriter, r *http.Request) {
	VideoTutorial := models.VideoTutorial{}

	VideoTutorials, err := VideoTutorial.GetAllVideoTutorials(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, VideoTutorials)
}

func (server *Server) GetVideoTutorial(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	VideoTutorial := models.VideoTutorial{}
	VideoTutorialGotten, err := VideoTutorial.GetVideoTutorialByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, VideoTutorialGotten)
}

func (server *Server) DeleteVideoTutorial(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	VideoTutorial := models.VideoTutorial{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = VideoTutorial.DeleteAVideoTutorial(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) UpdateVideoTutorial(w http.ResponseWriter, r *http.Request) {

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
	VideoTutorial := models.VideoTutorial{}
	err = json.Unmarshal(body, &VideoTutorial)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedVideoTutorial, err := VideoTutorial.UpdateVideoTutorial(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedVideoTutorial)
}
