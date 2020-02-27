package services

import (
	"encoding/json"
	"first/httpd/api/responses"
	"first/models"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (server *Server) Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	fmt.Println(user.Username)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.BeforeSave()
	user.SaveUser(server.DB)
}
