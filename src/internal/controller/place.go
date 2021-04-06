package controller

import (
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
	result := c.placeSvc.GetPlaceById(int64(ID))
	w.Write([]byte(fmt.Sprint(result.ID)))
	w.WriteHeader(http.StatusOK)
	return
}
