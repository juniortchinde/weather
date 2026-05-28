package main

import (
	"encoding/json"
	"net/http"
)

type App struct{ store *Store }

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func (a *App) listStations(w http.ResponseWriter, r *http.Request) {
	stations := a.store.All()
	if stations == nil {
		writeError(w, http.StatusBadRequest, "No stations found")
	} else {
		writeJSON(w, http.StatusOK, stations)
	}
}

func (a *App) getStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	station, ok := a.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "Station not found")
	} else {
		writeJSON(w, http.StatusOK, station)
	}
}

func (a *App) createStation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var station Station
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&station); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON : "+err.Error())
		return
	}
	if station.Id == "" {
		writeError(w, http.StatusBadRequest, "Invalid station ID")
		return
	}
	if a.store.Has(station.Id) {
		writeError(w, http.StatusConflict, "Station already exists")
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
		writeError(w, 400, "JSON invalide: "+err.Error())
		return
	}
	if st.Id != "" && st.Id != id {
		writeError(w, 400, "incohérence id body vs URL")
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
