package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (c *Controller) GetPlaceByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	ID, err := strconv.Atoi(strID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}
	result := c.placeSvc.GetPlaceById(r.Context(), int64(ID))
	w.Write([]byte(fmt.Sprint(result.ID)))
	w.WriteHeader(http.StatusOK)
	return
}

func (c *Controller) GetGymnasiumByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	ID, err := strconv.Atoi(strID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}

	result := c.locationSvc.GetGymnasiumByID(r.Context(), int64(ID))
	valByte, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshal"))
		return
	}

	w.Write(valByte)
	w.WriteHeader(http.StatusOK)
	return
}
