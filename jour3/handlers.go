package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type App struct{ store *Store }

type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, code, msg string) {
	writeJSON(w, status, ErrorResponse{Error: msg, Code: code})
}

func (a *App) listStations(w http.ResponseWriter, r *http.Request) {
	stations := a.store.All()
	if stations == nil {
		writeError(w, http.StatusBadRequest, "BAD_JSON", "No stations found")
	} else {
		writeJSON(w, http.StatusOK, stations)
	}
}

func (a *App) getStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	station, ok := a.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "NOT_FOUND", "Station not found")
	} else {
		writeJSON(w, http.StatusOK, station)
	}
}

// Crete new station
func (a *App) createStation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var station Station
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&station); err != nil {
		writeError(w, http.StatusBadRequest, "BAD_JSON", "Invalid JSON : "+err.Error())
		return
	}
	if station.Id == "" {
		writeError(w, http.StatusBadRequest, "BAD_JSON", "Invalid station ID")
		return
	}
	if a.store.Has(station.Id) {
		writeError(w, http.StatusConflict, "ID_TAKEN", "Station already exists")
		return
	}
	a.store.Put(station)
	writeJSON(w, http.StatusCreated, station)
}

func (a *App) updateStation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := r.PathValue("id")
	var st Station
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&st); err != nil {
		writeError(w, http.StatusBadRequest, "BAD_JSON", "JSON invalide: "+err.Error())
		return
	}
	if st.Id != "" && st.Id != id {
		writeError(w, http.StatusBadRequest, "BAD_JSON", "incohérence id body vs URL")
		return
	}

	st.Id = id
	created := !a.store.Has(id)
	a.store.Put(st)
	if created {
		writeJSON(w, 201, st)
		return
	}
	writeJSON(w, 200, st)
}

func (a *App) deleteStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if !a.store.Delete(id) {
		writeError(w, http.StatusNotFound, "NOT_FOUND",
			fmt.Sprintf("station %q introuvable", id))
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204 — pas de body
}
