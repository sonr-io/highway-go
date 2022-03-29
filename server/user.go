package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sonr-io/highway-go/logger"
	"github.com/sonr-io/highway-go/models"
)

// CreateUser adds a new user to the database
func (ws *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	if username == "" {
		jsonResponse(w, "No username specified", http.StatusBadRequest)
		return
	}
	if email == "" {
		jsonResponse(w, "No email specified", http.StatusBadRequest)
		return
	}
	user, err := ws.Ctrl.GetUserByUsername(email)
	if err != gorm.ErrRecordNotFound {
		log.Errorf("user already exists: %s", email)
		jsonResponse(w, user, http.StatusOK)
		return
	}
	u := models.User{
		Username:    email,
		DisplayName: username,
		Paid:        false,
		Icon:        models.PlaceholderUserIcon,
	}
	err = ws.Ctrl.PutUser(&u)
	if err != nil {
		jsonResponse(w, "Error Creating User", http.StatusInternalServerError)
		return
	}
	jsonResponse(w, u, http.StatusCreated)
}

// UserExists returns a boolean indicating if the user exists or not.
func (ws *Server) UserExists(w http.ResponseWriter, r *http.Request) {
	type existsResponse struct {
		Exists bool `json:"exists"`
	}
	vars := mux.Vars(r)
	//The trimmer
	username := vars["name"]
	if username[len(username)-4:] == ".snr" {
		username = username[:len(username)-4]
	}
	_, err := ws.Ctrl.GetUserByUsername(username)
	if err != nil {
		log.Errorf("user not found: %s: %s", username, err)
		jsonResponse(w, existsResponse{Exists: false}, http.StatusNotFound)
		return
	}
	jsonResponse(w, existsResponse{Exists: true}, http.StatusOK)
}
