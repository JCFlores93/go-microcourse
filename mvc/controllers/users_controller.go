package controllers

import (
	"encoding/json"
	"github.com/JCFlores93/go-microcourse/mvc/services"
	"log"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		log.Printf("about to process user_id %v", userIdParam)
		return
	}
	user, err := services.GetUser(userId)
	if err != nil {
		// handle the error
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}
	// return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)

}
