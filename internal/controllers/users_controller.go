package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Matheus002/go-microservice-rest/internal/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := (strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64))
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("Parameter error"))
		return
	}
	log.Printf("processing user_id %v", userId)
	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
