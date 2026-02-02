package handlers

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/Astra-max/backend-dev/go/0x00-go/pkg/server"
	"github.com/Astra-max/backend-dev/go/0x00-go/pkg/models"
)

var user models.UserLogins

func AuthUser(w http.ResponseWriter, req *http.Request) {
	data := map[string]string{"user": "welcome"}
	server.Json(w, 200, data)
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	newDb := new(models.DataBase)
	newDb.DB = append(newDb.DB, user)

	defer req.Body.Close()
}