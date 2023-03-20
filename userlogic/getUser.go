package userlogic

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
	"usersenaryo/databaseOP"
	"usersenaryo/models"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/get-user/"))
		data := databaseOP.Get(id)

		var user models.User
		user.ID = id
		hata := data.Scan(&user.Username, &user.FirstName, &user.Password)
		if hata != nil {
			log.Println("satır 23: ", hata)
			w.WriteHeader(http.StatusNotFound)
			notFound, _ := json.Marshal(map[string]string{"message": "kullanıcı bulunamadı"})
			fmt.Fprintf(w, string(notFound))
			return
		}
		veri, _ := json.Marshal(user)
		fmt.Fprintf(w, string(veri))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed) // metot uygun değilse 405 kodu dön
	}

}
