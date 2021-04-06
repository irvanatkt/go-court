package controller

import "net/http"

func (c *Controller) GetPlaceByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	w.WriteHeader(http.StatusOK)
}
