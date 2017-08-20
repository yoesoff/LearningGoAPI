package app

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	. "gitlab.com/mhyusufibrahim/teahrm/app/models/user"
)

// Request Handler to get bunch of users
func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	users, err := GetUsers(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

// Request Handler to create a single user
func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var u User
	var decoder = json.NewDecoder(r.Body)

	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := u.CreateUser(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err2 := u.GetUser(a.DB); err2 != nil {
		switch err2 {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err2.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusCreated, u)
}

// Request Handler to get a single user
func (a *App) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	u := User{ID: id}
	if err := u.GetUser(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, u)
}

// Request Handler to update a single user
func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var u User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}

	u.ID = id

	if err := u.UpdateUser(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, u)
}

// Request Handler to delete a single user
func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	u := User{ID: id}
	if err := u.DeleteUser(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// Get existing token by email and password (Login)
func (a *App) getApiToken(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var auth Auth
	var decoder = json.NewDecoder(r.Body)

	if err := decoder.Decode(&auth); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	auth.GetApiToken(a.DB)

	respondWithJSON(w, http.StatusOK, auth)
}

// Update token by existing token (Logout)
func (a *App) changeApiToken(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var auth Auth
	var decoder = json.NewDecoder(r.Body)

	if err := decoder.Decode(&auth); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := auth.ChangeApiToken(a.DB); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID or Token")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})

}
